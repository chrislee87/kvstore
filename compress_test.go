package kvstore

import (
	"reflect"
	"testing"
)

var poem string = `
O how much more doth beauty beauteous seem, 
By that sweet ornament which truth doth give!
The rose looks fair, but fairer we it deem 
For that sweet odour which doth in it live. 
The canker-blooms have full as deep a dye 
As the perfumed tincture of the roses, 
Hang on such thorns and play as wantonly 
When summer's breath their masked buds discloses:
But, for their virtue only is their show, 
They live unwoo'd and unrespected fade, 
Die to themselves. Sweet roses do not so; 
Of their sweet deaths are sweetest odours made:
And so of you, beauteous and lovely youth,
When that shall fade, my verse distills your truth.
`

func TestGzip(t *testing.T) {
	bs := []byte(poem)
	b1 := GzipCompress(bs)
	b2 := GzipUnCompress(b1)

	if !reflect.DeepEqual(bs, b2) {
		t.Errorf("After GzipCompress and GzipUncompress, slice of bytes should be the same.")
	}
}

func TestSnappy(t *testing.T) {
	bs := []byte(poem)
	b1 := SnappyCompress(bs)
	b2 := SnappyUnCompress(b1)

	if !reflect.DeepEqual(bs, b2) {
		t.Errorf("After SnappyCompress and SnappyUncompress, slice of bytes should be the same.")
	}
}
