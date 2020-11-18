package service

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"ditto.co.jp/grpc/pb"
	"ditto.co.jp/grpc/w32"
	"golang.org/x/sys/windows"
)

//BrowseForFolder -
func (s *LocalService) BrowseForFolder(ctx context.Context, msg *pb.BrowseInfo) (*pb.Folder, error) {
	d := w32.DirectoryBuilder{
		Title:          msg.Title,
		InitialDirPath: msg.Initialdir,
	}
	path, err := d.Browse()
	if err != nil {
		return nil, err
	}
	result := &pb.Folder{
		Path: path,
	}

	return result, nil
}

//Alive -
func (s *LocalService) Alive(ctx context.Context, msg *pb.Empty) (*pb.AliveResponse, error) {
	homedir := ""
	t, err := windows.OpenCurrentProcessToken()
	if err == nil {
		dir, err := t.GetUserProfileDirectory()
		if err == nil {
			homedir = dir
		}
	}
	if homedir == "" {
		dir, err := dirWindows()
		if err != nil {
			dir = "~"
		} else {
			dir = filepath.ToSlash(homedir)
		}
		homedir = dir
	}

	pid := os.Getpid()
	result := &pb.AliveResponse{
		Processid: int32(pid),
		Homedir:   homedir,
		Accesskey: os.Getenv("AWS_ACCESS_KEY_ID"),
		Secretkey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:    os.Getenv("AWS_DEFAULT_REGION"),
	}
	return result, nil
}

func dirWindows() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// Prefer standard environment variable USERPROFILE
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
}
