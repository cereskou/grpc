package w32

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

//SHBrowseForFolder flags
const (
	BIF_RETURNONLYFSDIRS    = 0x00000001
	BIF_DONTGOBELOWDOMAIN   = 0x00000002
	BIF_STATUSTEXT          = 0x00000004
	BIF_RETURNFSANCESTORS   = 0x00000008
	BIF_EDITBOX             = 0x00000010
	BIF_VALIDATE            = 0x00000020
	BIF_NEWDIALOGSTYLE      = 0x00000040
	BIF_BROWSEINCLUDEURLS   = 0x00000080
	BIF_USENEWUI            = BIF_EDITBOX | BIF_NEWDIALOGSTYLE
	BIF_UAHINT              = 0x00000100
	BIF_NONEWFOLDERBUTTON   = 0x00000200
	BIF_NOTRANSLATETARGETS  = 0x00000400
	BIF_BROWSEFORCOMPUTER   = 0x00001000
	BIF_BROWSEFORPRINTER    = 0x00002000
	BIF_BROWSEINCLUDEFILES  = 0x00004000
	BIF_SHAREABLE           = 0x00008000
	BIF_BROWSEFILEJUNCTIONS = 0x00010000
)

var browseFolderCallbackPtr uintptr

//DirectoryBuilder -
type DirectoryBuilder struct {
	Title          string
	InitialDirPath string
}

// Browse -
func (d *DirectoryBuilder) Browse() (string, error) {
	if hr := win.OleInitialize(); hr != win.S_OK && hr != win.S_FALSE {
		return "", fmt.Errorf("OleInitialize error: %v", hr)
	}
	defer win.OleUninitialize()

	bi := &win.BROWSEINFO{
		HwndOwner: win.GetForegroundWindow(),
		UlFlags:   BIF_RETURNONLYFSDIRS | BIF_NEWDIALOGSTYLE,
	}
	if d.InitialDirPath != "" {
		callback := func(hwnd win.HWND, msg uint32, lp, wp uintptr) uintptr {
			const BFFM_INITIALIZED = 1
			const BFFM_SELCHANGED = 2
			const BFFM_SETSELECTION = win.WM_USER + 103

			if msg == BFFM_INITIALIZED { // ini send a path to select
				upath := strings.Replace(d.InitialDirPath, "/", "\\", -1)
				win.SendMessage(hwnd, BFFM_SETSELECTION, win.TRUE, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(upath))))
			}

			if msg == BFFM_SELCHANGED {
				_, err := pathFromPIDL(lp)
				var enabled uintptr
				if err == nil {
					enabled = 1
				}
				const BFFM_ENABLEOK = win.WM_USER + 101

				win.SendMessage(hwnd, BFFM_ENABLEOK, 0, enabled)
			}

			return 0
		}
		bi.Lpfn = syscall.NewCallback(callback)
	}
	if d.Title != "" {
		bi.LpszTitle = syscall.StringToUTF16Ptr(d.Title)
	}

	pidl := win.SHBrowseForFolder(bi)
	if pidl == 0 {
		return "", nil
	}
	defer win.CoTaskMemFree(pidl)

	return pathFromPIDL(pidl)
}

func pathFromPIDL(pidl uintptr) (string, error) {
	var path [win.MAX_PATH]uint16
	if !win.SHGetPathFromIDList(pidl, &path[0]) {
		return "", fmt.Errorf("SHGetPathFromIDList failed")
	}

	return syscall.UTF16ToString(path[:]), nil
}
