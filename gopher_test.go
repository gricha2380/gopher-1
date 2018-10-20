package gopher

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

func TestTalk(t *testing.T) {
	testScream := new(bool)
	*testScream = true

	tests := []struct {
		msg    string
		scream *bool
		want   string
	}{
		{
			msg:    "Concurrency is not parallelism",
			scream: new(bool),
			want:   "Concurrency is not parallelism",
		},
		{
			want:   "WHAT AM I SUPPOSED TO SAY?",
			scream: testScream,
		},
		{
			want:   "What am I supposed to say?",
			scream: new(bool),
		},
		{
			msg:    "Make the zero value useful",
			scream: new(bool),
			want:   "Make the zero value useful",
		},
		{
			msg:    "Don't panic!",
			scream: testScream,
			want:   "DON'T PANIC!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			out := &bytes.Buffer{}
			Talk(tt.msg, tt.scream, out)

			gotOut := out.String()
			re := regexp.MustCompile(tt.want)

			if match := re.FindString(gotOut); match == "" {
				t.Errorf("Talk() = %s, want to include %q", gotOut, tt.want)
			}
		})
	}
}

func BenchmarkHTalk(b *testing.B) {
	out := &bytes.Buffer{}

	for i := 0; i < b.N; i++ {
		out.Reset()
		Talk("", new(bool), out)
	}
}

func ExampleTalk() {
	Talk("Hello, there!", new(bool), os.Stdout)
	// Output:
	// ~~~~~~~~~~~~~~~
	// / Hello, there! /
	//  ~~~~~~~~~~~~~~~
	//     ###
	//       ##
	//         #
	// 	                        dhyysoo++++ooossyhd
	// 	                Ndyo/:#######################:+sd
	// 	      NN/:###ds/#######################/+////+/###+y#######
	// 	  Ny/:###/hy:#:+/::###::/+:#########+/:"  .:#  :+/##:oo######
	// 	 h:##//:+o:#:+:  +hdy.    #+:#####++"     s####" :+###:sN +###:
	// 	 ###d   /##:o   :#####      o:###//       h#s#N.  :/####+Ny####
	// 	d###sN :###s"    "sd+d:     "o###s         ://"    s#####++###/
	// 	 y:##s/####s                "s###s"               "s######s:/
	// 	  Nhso#####//               +:###:o"             "o:######:
	// 	    N#######+/            "+/:+sss+o/"         ./+#########
	// 	    y########:+/:.    "#://#+N     N:+//:::::/+/###########/
	// 	    +############:/+++/:##ooodNNN dsoo+#####################
	// 	    /####################s/:::::::::::/s####################h
	// 	    :####################+o//+++s/++//oo####################s
	// 	    :######################o/  ##  ++/:#####################o
	// 	    /######################:+  ::  :/#######################+
	// 	    +#######################+//++//+########################+
}
