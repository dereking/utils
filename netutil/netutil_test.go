package netutil

import (
	"github.com/Dereking/utils/osutil"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	tmp := osutil.TempFile("downlaod")
	t.Log("temp path=", tmp)
	Download("http://baidu.com", tmp)
	ok, err := osutil.PathExists(tmp)
	if err != nil {
		t.Error("download err", err)
	} else {
		t.Log("download", ok, tmp)
		os.Remove(tmp)
	}
}
