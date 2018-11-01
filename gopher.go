// Package gopher writes a simple and funny ASCII gopher with a dialog.
package gopher

import (
	"fmt"
	"io"
	"strings"
)

// Talk writes a funny ASCII gopher with a dialog at the specified writer.
// If the upper *bool its true it writes the message in uppercase.
func Talk(msg string, upper *bool, out io.Writer) {
	// ASCII art of a wild gopher.
	ASCIIart := `
			        dhyysoo++++ooossyhd
	                Ndyo/:#######################:+sd
	      NN/:###ds/#######################/+////+/###+y#######
	  Ny/:###/hy:#:+/::###::/+:#########+/:"  .:#  :+/##:oo######
	 h:##//:+o:#:+:  +hdy.    #+:#####++"     s####" :+###:sN +###:
	 ###d   /##:o   :#####      o:###//       h#s#N.  :/####+Ny####
	d###sN :###s"    "sd+d:     "o###s         ://"    s#####++###/
	 y:##s/####s                "s###s"               "s######s:/
	  Nhso#####//               +:###:o"             "o:######:
	    N#######+/            "+/:+sss+o/"         ./+#########
	    y########:+/:.    "#://#+N     N:+//:::::/+/###########/
	    +############:/+++/:##ooodNNN dsoo+#####################
	    /####################s/:::::::::::/s####################h
	    :####################+o//+++s/++//oo####################s
	    :######################o/  ##  ++/:#####################o
	    /######################:+  ::  :/#######################+
	    +#######################+//++//+########################+`

	dialog := `
	###
	  ##
	    #`

	ml := len(msg)

	if ml == 0 {
		msg = "What am I supposed to say?"
		ml = len(msg)
	}

	if *upper {
		msg = strings.ToUpper(msg)
	}

	io.WriteString(out, fmt.Sprintf(" %s\n/ %s /\n %[1]s", strings.Repeat("~", ml+2), msg))
	io.WriteString(out, fmt.Sprintf("%s%s", strings.Repeat(" ", ml/3), dialog))
	io.WriteString(out, ASCIIart+"\n")
}
