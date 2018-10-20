// Package gopher generates a simple and funny ASCII gopher with a message.
package gopher

import (
	"fmt"
	"io"
	"strings"
)

// Talk writes a dialog at the specified writer.
// If the upper *bool its true it writes the message in uppercase.
func Talk(msg string, upper *bool, out io.Writer) {
	// ASCII art of the Go project's pet.
	ASCIIart := `	                        dhyysoo++++ooossyhd
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
	    +#######################+//++//+########################+
		`

	ml := len(msg)

	if ml <= 0 {
		msg = "What am I supposed to say?"
		ml = len(msg)
	}

	if *upper {
		msg = strings.ToUpper(msg)
	}

	io.WriteString(out, fmt.Sprintf(" %s\n/ %s /\n %[1]s\n", strings.Repeat("~", ml+2), msg))
	for i := 0; i < 3; i++ {
		io.WriteString(out, fmt.Sprintf("%s%s\n", strings.Repeat(" ", (ml/3)+(i*2)), strings.Repeat("#", 3-i)))
	}

	io.WriteString(out, ASCIIart)
}
