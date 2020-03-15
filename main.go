package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/wargarblgarbl/libgosubs/srt"
	webvtt "github.com/wargarblgarbl/libgosubs/wvtt"
	subconvfunc "github.com/wargarblgarbl/subconvfunc"
)

var inpath string
var outpath string

var subtitles []*srt.Subtitle

func main() {
	v := &srt.SubRip{}
	if len(os.Args) > 1 {
		inpath = os.Args[1]
	} else {
		fmt.Println("Usage : " + os.Args[0] + " input_file_name output_file_name")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		outpath = os.Args[2]
	} else {
		outpath = strings.Replace(inpath, ".vtt", ".srt", -1)
	}
	loadedvtt, err := webvtt.ParseWebVtt(inpath)
	if err != nil {
		panic(err)
	}
	for a, i := range loadedvtt.Subtitle.Content {
		z := srt.CreateSubtitle(a+1, subconvfunc.TtmlToSrtTimecode(i.Start), subconvfunc.TtmlToSrtTimecode(i.End), i.Line)
		v.Subtitle.Content = append(v.Subtitle.Content, *z)
	}
	srt.WriteSrt(v, outpath)
}
