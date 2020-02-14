package main

import(
	"os"
  "path/filepath"
  //"github.com/facebookgo/symwalk"

  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "log"
)

 // DirectoryIterator iterates through files and directories to be uploaded                                          
// to S3.                                                                                                               
type DirectoryIterator struct {                             
  filePaths []string           
  bucket    string                      
  next      struct {                        
    path string                        
    f    *os.File
  }                                   
  err error                                       
}                               
// NewDirectoryIterator creates and returns a new BatchUploadIterator                                                
func NewDirectoryIterator(bucket, dir string) s3manager.BatchUploadIterator {
  log.Println(dir)
  paths := []string{}
  f,_ := filepath.Abs(dir)
  log.Println("f", f)
  filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {                                             
    // We care only about files, not directories  
    log.Println("INFO: ", info)
    log.Println("ERROR: ", err)
                                                                    
    if !info.IsDir() {
      log.Println("PATH: ")
      log.Println(path)
			paths = append(paths, path)                                        
    }           
    return nil                                       
  })                          
  return &DirectoryIterator{
		filePaths: paths,
		bucket:    bucket,                                        
  }                                                                                                                     
}                  

// Next opens the next file and stops iteration if it fails to open                                             
// a file.                                                                                                              
func (iter *DirectoryIterator) Next() bool {                        
  if len(iter.filePaths) == 0 {                                        
    iter.next.f = nil             
    return false                       
	}
	                        
  f, err := os.Open(iter.filePaths[0])                                        
  iter.err = err  
  iter.next.f = f                           
  iter.next.path = iter.filePaths[0]                        
  iter.filePaths = iter.filePaths[1:]      
  return true && iter.Err() == nil     
}      

// Err returns an error that was set during opening the file
func (iter *DirectoryIterator) Err() error {
  return iter.err
}

// UploadObject returns a BatchUploadObject and sets the After field to                                              
// close the file.                                                                                                      
func (iter *DirectoryIterator) UploadObject() s3manager.BatchUploadObject {
	f := iter.next.f
	return s3manager.BatchUploadObject{                                        
    Object: &s3manager.UploadInput{       
      Bucket: &iter.bucket,           
      Key:    &iter.next.path,                   
      Body:   f,                
		},
		
	// After was introduced in version 1.10.7
    After: func() error {                                        
      return f.Close()                     
    },                      
  }                                    
}
