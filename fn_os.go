package xtemplate

import (
	"os"
	"time"

	"github.com/Eun/xtemplate/funcs"
)

// OS provides access to functions in the os package.
type OS rootContext

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
//
// Example:
//
//	{{ os.Chdir "/tmp" }}
func (ctx OS) Chdir(dir string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSChdir]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSChdir}
	}
	return os.Chdir(dir)
}

// Chmod changes the mode of the named file to mode.
// If the file is a symbolic link, it changes the mode of the link's target.
//
// Example:
//
//	{{ os.Chmod "file.txt" 0644 }}
func (ctx OS) Chmod(name string, mode os.FileMode) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSChmod]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSChmod}
	}
	return os.Chmod(name, mode)
}

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
//
// Example:
//
//	{{ os.Chown "file.txt" 1000 1000 }}
func (ctx OS) Chown(name string, uid, gid int) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSChown]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSChown}
	}
	return os.Chown(name, uid, gid)
}

// Chtimes changes the access and modification times of the named file,
// similar to the Unix utime() or utimes() functions.
//
// Example:
//
//	{{ os.Chtimes "file.txt" .AccessTime .ModTime }}
func (ctx OS) Chtimes(name string, atime time.Time, mtime time.Time) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSChtimes]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSChtimes}
	}
	return os.Chtimes(name, atime, mtime)
}

// Clearenv deletes all environment variables.
//
// Example:
//
//	{{ os.Clearenv }}
func (ctx OS) Clearenv() error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSClearenv]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSClearenv}
	}
	os.Clearenv()
	return nil
}

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
//
// Example:
//
//	{{ os.Environ }}
func (ctx OS) Environ() ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSEnviron]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.OSEnviron}
	}
	return os.Environ(), nil
}

// Executable returns the path name for the executable that started
// the current process.
//
// Example:
//
//	{{ os.Executable }}
func (ctx OS) Executable() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSExecutable]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSExecutable}
	}
	return os.Executable()
}

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
//
// Example:
//
//	{{ os.Exit 0 }}
func (ctx OS) Exit(code int) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSExit]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSExit}
	}
	os.Exit(code)
	return nil
}

// Expand replaces ${var} or $var in the string based on the mapping function.
// For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).
//
// Example:
//
//	{{ os.Expand "$HOME/file" .MappingFunc }}
func (ctx OS) Expand(s string, mapping func(string) string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSExpand]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSExpand}
	}
	return os.Expand(s, mapping), nil
}

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables.
//
// Example:
//
//	{{ os.ExpandEnv "$HOME/file" }}
func (ctx OS) ExpandEnv(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSExpandEnv]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSExpandEnv}
	}
	return os.ExpandEnv(s), nil
}

// Getegid returns the numeric effective group id of the caller.
// On Windows, it returns -1.
//
// Example:
//
//	{{ os.Getegid }}
func (ctx OS) Getegid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetegid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetegid}
	}
	return os.Getegid(), nil
}

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
//
// Example:
//
//	{{ os.Getenv "HOME" }}
func (ctx OS) Getenv(key string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetenv]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSGetenv}
	}
	return os.Getenv(key), nil
}

// Geteuid returns the numeric effective user id of the caller.
// On Windows, it returns -1.
//
// Example:
//
//	{{ os.Geteuid }}
func (ctx OS) Geteuid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGeteuid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGeteuid}
	}
	return os.Geteuid(), nil
}

// Getgid returns the numeric group id of the caller.
// On Windows, it returns -1.
//
// Example:
//
//	{{ os.Getgid }}
func (ctx OS) Getgid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetgid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetgid}
	}
	return os.Getgid(), nil
}

// Getgroups returns a list of the numeric ids of groups that the caller belongs to.
//
// Example:
//
//	{{ os.Getgroups }}
func (ctx OS) Getgroups() ([]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetgroups]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.OSGetgroups}
	}
	return os.Getgroups()
}

// Getpagesize returns the underlying system's memory page size.
//
// Example:
//
//	{{ os.Getpagesize }}
func (ctx OS) Getpagesize() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetpagesize]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetpagesize}
	}
	return os.Getpagesize(), nil
}

// Getpid returns the process id of the caller.
//
// Example:
//
//	{{ os.Getpid }}
func (ctx OS) Getpid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetpid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetpid}
	}
	return os.Getpid(), nil
}

// Getppid returns the process id of the caller's parent.
//
// Example:
//
//	{{ os.Getppid }}
func (ctx OS) Getppid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetppid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetppid}
	}
	return os.Getppid(), nil
}

// Getuid returns the numeric user id of the caller.
// On Windows, it returns -1.
//
// Example:
//
//	{{ os.Getuid }}
func (ctx OS) Getuid() (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetuid]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.OSGetuid}
	}
	return os.Getuid(), nil
}

// Getwd returns a rooted path name corresponding to the
// current directory.
//
// Example:
//
//	{{ os.Getwd }}
func (ctx OS) Getwd() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSGetwd]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSGetwd}
	}
	return os.Getwd()
}

// Hostname returns the host name reported by the kernel.
//
// Example:
//
//	{{ os.Hostname }}
func (ctx OS) Hostname() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSHostname]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSHostname}
	}
	return os.Hostname()
}

// IsExist returns a boolean indicating whether the error is known to report
// that a file or directory already exists.
//
// Example:
//
//	{{ os.IsExist .Error }}
func (ctx OS) IsExist(err error) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSIsExist]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSIsExist}
	}
	return os.IsExist(err), nil
}

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist.
//
// Example:
//
//	{{ os.IsNotExist .Error }}
func (ctx OS) IsNotExist(err error) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSIsNotExist]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSIsNotExist}
	}
	return os.IsNotExist(err), nil
}

// IsPathSeparator reports whether c is a directory separator character.
//
// Example:
//
//	{{ os.IsPathSeparator 47 }}
func (ctx OS) IsPathSeparator(c uint8) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSIsPathSeparator]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSIsPathSeparator}
	}
	return os.IsPathSeparator(c), nil
}

// IsPermission returns a boolean indicating whether the error is known to
// report that permission is denied.
//
// Example:
//
//	{{ os.IsPermission .Error }}
func (ctx OS) IsPermission(err error) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSIsPermission]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSIsPermission}
	}
	return os.IsPermission(err), nil
}

// IsTimeout returns a boolean indicating whether the error is known
// to report that a timeout occurred.
//
// Example:
//
//	{{ os.IsTimeout .Error }}
func (ctx OS) IsTimeout(err error) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSIsTimeout]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSIsTimeout}
	}
	return os.IsTimeout(err), nil
}

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
//
// Example:
//
//	{{ os.Lchown "file.txt" 1000 1000 }}
func (ctx OS) Lchown(name string, uid, gid int) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSLchown]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSLchown}
	}
	return os.Lchown(name, uid, gid)
}

// Link creates newname as a hard link to the oldname file.
//
// Example:
//
//	{{ os.Link "oldfile" "newfile" }}
func (ctx OS) Link(oldname, newname string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSLink]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSLink}
	}
	return os.Link(oldname, newname)
}

// LookupEnv retrieves the value of the environment variable named
// by the key. If the variable is present in the environment the
// value (which may be empty) is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
//
// Example:
//
//	{{ os.LookupEnv "HOME" }}
func (ctx OS) LookupEnv(key string) (string, bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSLookupEnv]; !ok {
		return "", false, &FuncNotAllowedError{Func: funcs.OSLookupEnv}
	}
	value, found := os.LookupEnv(key)
	return value, found, nil
}

// Mkdir creates a new directory with the specified name and permission
// bits (before umask).
//
// Example:
//
//	{{ os.Mkdir "newdir" 0755 }}
func (ctx OS) Mkdir(name string, perm os.FileMode) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSMkdir]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSMkdir}
	}
	return os.Mkdir(name, perm)
}

// MkdirAll creates a directory named path, along with any necessary
// parents, and returns nil, or else returns an error.
//
// Example:
//
//	{{ os.MkdirAll "path/to/dir" 0755 }}
func (ctx OS) MkdirAll(path string, perm os.FileMode) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSMkdirAll]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSMkdirAll}
	}
	return os.MkdirAll(path, perm)
}

// MkdirTemp creates a new temporary directory in the directory dir
// and returns the pathname of the new directory.
//
// Example:
//
//	{{ os.MkdirTemp "/tmp" "pattern" }}
func (ctx OS) MkdirTemp(dir, pattern string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSMkdirTemp]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSMkdirTemp}
	}
	return os.MkdirTemp(dir, pattern)
}

// NewSyscallError returns, as an error, a new SyscallError
// with the given system call name and error details.
//
// Example:
//
//	{{ os.NewSyscallError "open" .Error }}
func (ctx OS) NewSyscallError(syscall string, err error) (error, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSNewSyscallError]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.OSNewSyscallError}
	}
	return os.NewSyscallError(syscall, err), nil
}

// Pipe returns a connected pair of Files; reads from r return bytes written to w.
//
// Example:
//
//	{{ os.Pipe }}
func (ctx OS) Pipe() (*os.File, *os.File, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSPipe]; !ok {
		return nil, nil, &FuncNotAllowedError{Func: funcs.OSPipe}
	}
	return os.Pipe()
}

// ReadFile reads the named file and returns the contents.
//
// Example:
//
//	{{ os.ReadFile "file.txt" }}
func (ctx OS) ReadFile(name string) ([]byte, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSReadFile]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.OSReadFile}
	}
	return os.ReadFile(name) //nolint:gosec // G304: allowed function
}

// Readlink returns the destination of the named symbolic link.
//
// Example:
//
//	{{ os.Readlink "symlink" }}
func (ctx OS) Readlink(name string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSReadlink]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSReadlink}
	}
	return os.Readlink(name)
}

// Remove removes the named file or (empty) directory.
//
// Example:
//
//	{{ os.Remove "file.txt" }}
func (ctx OS) Remove(name string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSRemove]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSRemove}
	}
	return os.Remove(name)
}

// RemoveAll removes path and any children it contains.
// It removes everything it can but returns the first error it encounters.
//
// Example:
//
//	{{ os.RemoveAll "path/to/dir" }}
func (ctx OS) RemoveAll(path string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSRemoveAll]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSRemoveAll}
	}
	return os.RemoveAll(path)
}

// Rename renames (moves) oldpath to newpath.
// If newpath already exists and is not a directory, Rename replaces it.
//
// Example:
//
//	{{ os.Rename "oldname" "newname" }}
func (ctx OS) Rename(oldpath, newpath string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSRename]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSRename}
	}
	return os.Rename(oldpath, newpath)
}

// SameFile reports whether fi1 and fi2 describe the same file.
//
// Example:
//
//	{{ os.SameFile .FileInfo1 .FileInfo2 }}
func (ctx OS) SameFile(fi1, fi2 os.FileInfo) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSSameFile]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.OSSameFile}
	}
	return os.SameFile(fi1, fi2), nil
}

// Setenv sets the value of the environment variable named by the key.
//
// Example:
//
//	{{ os.Setenv "KEY" "value" }}
func (ctx OS) Setenv(key, value string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSSetenv]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSSetenv}
	}
	return os.Setenv(key, value)
}

// Symlink creates newname as a symbolic link to oldname.
//
// Example:
//
//	{{ os.Symlink "oldname" "newname" }}
func (ctx OS) Symlink(oldname, newname string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSSymlink]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSSymlink}
	}
	return os.Symlink(oldname, newname)
}

// TempDir returns the default directory to use for temporary files.
//
// Example:
//
//	{{ os.TempDir }}
func (ctx OS) TempDir() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSTempDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSTempDir}
	}
	return os.TempDir(), nil
}

// Truncate changes the size of the named file.
//
// Example:
//
//	{{ os.Truncate "file.txt" 100 }}
func (ctx OS) Truncate(name string, size int64) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSTruncate]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSTruncate}
	}
	return os.Truncate(name, size)
}

// Unsetenv unsets a single environment variable.
//
// Example:
//
//	{{ os.Unsetenv "KEY" }}
func (ctx OS) Unsetenv(key string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSUnsetenv]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSUnsetenv}
	}
	return os.Unsetenv(key)
}

// UserCacheDir returns the default root directory to use for user-specific cached data.
//
// Example:
//
//	{{ os.UserCacheDir }}
func (ctx OS) UserCacheDir() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSUserCacheDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSUserCacheDir}
	}
	return os.UserCacheDir()
}

// UserConfigDir returns the default root directory to use for user-specific configuration data.
//
// Example:
//
//	{{ os.UserConfigDir }}
func (ctx OS) UserConfigDir() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSUserConfigDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSUserConfigDir}
	}
	return os.UserConfigDir()
}

// UserHomeDir returns the current user's home directory.
//
// Example:
//
//	{{ os.UserHomeDir }}
func (ctx OS) UserHomeDir() (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.OSUserHomeDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.OSUserHomeDir}
	}
	return os.UserHomeDir()
}

// WriteFile writes data to the named file, creating it if necessary.
//
// Example:
//
//	{{ os.WriteFile "file.txt" .Data 0644 }}
func (ctx OS) WriteFile(name string, data []byte, perm os.FileMode) error {
	if _, ok := ctx.allowedFunctionSet[funcs.OSWriteFile]; !ok {
		return &FuncNotAllowedError{Func: funcs.OSWriteFile}
	}
	return os.WriteFile(name, data, perm)
}
