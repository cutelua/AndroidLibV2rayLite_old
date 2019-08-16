package shippedBinarys

import (	
	"log"
	"os"
	//"strconv"
	 "fmt"
    "io"
	"github.com/2dust/AndroidLibV2rayLite/CoreI"
)

type FirstRun struct {
	Status *CoreI.Status
}

func (v *FirstRun) checkIfRcExist() error {
	datadir := v.Status.GetDataDir()
	/*
	if _, err := os.Stat(datadir + strconv.Itoa(CoreI.CheckVersion())); !os.IsNotExist(err) {
		log.Println("file exists")
		return nil
	}
	*/
	log.Println("checkIfRcExist...")
	
	/*
	IndepDir, err := AssetDir("ArchIndep")
	log.Println(IndepDir)
	if err != nil {
		return err
	}
	for _, fn := range IndepDir {
		log.Println(datadir+"ArchIndep/"+fn)
		
		//err := RestoreAsset(datadir, "ArchIndep/"+fn)
		//log.Println(err)
		
		//GrantPremission
		//os.Chmod(datadir+"ArchIndep/"+fn, 0700)
		//log.Println(os.Remove(datadir + fn))	
		log.Println(CopyFile(datadir+"ArchIndep/"+fn, datadir + fn))			
	}
	*/
	
	DepDir, err := AssetDir("ArchDep")
	log.Println(DepDir)
	if err != nil {
		return err
	}
	for _, fn := range DepDir {
		DepDir2, err := AssetDir("ArchDep/" + fn)
		log.Println("ArchDep/" + fn)
		if err != nil {
			return err
		}
		for _, FND := range DepDir2 {
			log.Println(datadir+"ArchDep/"+fn+"/"+FND)
			
			//RestoreAsset(datadir, "ArchDep/"+fn+"/"+FND)
			//os.Chmod(datadir+"ArchDep/"+fn+"/"+FND, 0700)
			//log.Println(os.Remove(datadir + FND))	
			log.Println(CopyFile(datadir+"ArchDep/"+fn+"/"+FND, datadir+FND))			
		}
	}
	//s, _ := os.Create(datadir + strconv.Itoa(CoreI.CheckVersion()))
	//s.Close()

	return nil
}

func (v *FirstRun) CheckAndExport() error {
	return v.checkIfRcExist()
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherise, attempt to create a hard link
// between the two files. If that fail, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
    sfi, err := os.Stat(src)
    if err != nil {
        return
    }
    if !sfi.Mode().IsRegular() {
        // cannot copy non-regular files (e.g., directories,
        // symlinks, devices, etc.)
        return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
    }
    dfi, err := os.Stat(dst)
    if err != nil {
        if !os.IsNotExist(err) {
            return
        }
    } else {
        if !(dfi.Mode().IsRegular()) {
            return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
        }
        if os.SameFile(sfi, dfi) {
			log.Println("SameFile")
            return
        }
    }
	log.Println("start Symlink file")	
    if err = os.Symlink(src, dst); err == nil {	
		log.Println("Symlink file success")	
		os.Chmod(dst, 0700)
        return
    }
    err = copyFileContents(src, dst)
    return
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
    in, err := os.Open(src)
    if err != nil {
        return
    }
    defer in.Close()
    out, err := os.Create(dst)
    if err != nil {
        return
    }
    defer func() {
        cerr := out.Close()
        if err == nil {
            err = cerr
        }
    }()
    if _, err = io.Copy(out, in); err != nil {
        return
    }
    err = out.Sync()
	log.Println("copy File Contents")	
	os.Chmod(dst, 0700)
    return
}
