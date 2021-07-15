/*
Copyright (C) 2021 Colin Hughes

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"flag"
)

var Data TextData
var KPS []float64

//var SameKeyKPS []float64

func init() {
}

func main() {
	flag.StringVar(&ImproveFlag, "improve", "", "if set, decides which layout to improve")
	flag.BoolVar(&StaggerFlag, "stagger", false, "if true, calculates distance for row-stagger form factor")
	flag.BoolVar(&SlideFlag, "slide", false, "if true, ignores slideable sfbs")
	flag.BoolVar(&DynamicFlag, "dynamic", false, "")
	flag.Parse()
	origargs := os.Args[1:]
	var args []string
	for _, v := range origargs {
		if string(v[0]) != "-" {
			args = append(args, v)
		}
	}
	GeneratePositions()
	KPS = []float64{1.0, 4.0, 4.8, 5.7, 5.7, 4.8, 4.0, 1.0}
	//KPS = []float64{6, 16, 26.5, 40.36, 40.36, 26.5, 16, 6}
	//KPS = []float64{1, 3, 6, 8, 8, 6, 3, 1}
	//KPS = []float64{1, 1, 1, 1, 1, 1, 1, 1}

	Layouts = make(map[string]Layout)

	Layouts["qwerty"] = NewLayout("QWERTY", "qwertyuiopasdfghjkl;zxcvbnm,./")
	//Layouts["azerty"] = "azertyuiopqsdfghjklmwxcvbn',./"
	Layouts["dvorak"] = NewLayout("Dvorak", "',.pyfgcrlaoeuidhtns;qjkxbmwvz")
	Layouts["colemak"] = NewLayout("Colemak", "qwfpgjluy;arstdhneiozxcvbkm,./")
	Layouts["colemak dh"] = NewLayout("Colemak DH", "qwfpbjluy;arstgmneiozxcdvkh,./")
	// Layouts["funny colemak dh"] = "qwfpbjkuy;arstgmneiozxcdvlh,./"

	Layouts["colemaq"] = NewLayout("ColemaQ", ";wfpbjluyqarstgmneiozxcdkvh/.,")
	Layouts["colemaq-f"] = NewLayout("ColemaQ-f", ";wgpbjluyqarstfmneiozxcdkvh/.,")
	Layouts["colemak f"] = NewLayout("Colemak F", "qwgpbjluy;arstfmneiozxcdvkh,./")
	Layouts["colemak qi"] = NewLayout("Colemak Qi", "qlwmkjfuy'arstgpneiozxcdvbh,./")
	Layouts["colemak qix"] = NewLayout("Colemak Qi;x", ";lcmkjfuyqarstgpneiozxwdvbh/.,")
	// Layouts["NESO"] = "qylmkjfuc;airtgpnesoz.wdvbh/x,"
	// Layouts["NESO 2"] = "qylwvjfuc;airtgpneso.zkdmbh,x/"
	// "qulmkzbocyairtgpnesh.,wdjvf;x/"
	Layouts["isrt"] = NewLayout("ISRT", "yclmkzfu,'isrtgpneaoqvwdjbh/.x")
	// Layouts["hands down"] = "qchpvkyoj/rsntgwueiaxmldbzf',."
	Layouts["norman"] = NewLayout("Norman","qwdfkjurl;asetgyniohzxcvbpm,./")
	Layouts["mtgap"] = NewLayout("MTGAP", "ypoujkdlcwinea,mhtsrqz/.;bfgvx")
	Layouts["mtgap 2.0"] = NewLayout("MTGAP 2.0", ",fhdkjcul.oantgmseriqxbpzyw'v;")
	Layouts["sind"] = NewLayout("SIND", "y,hwfqkouxsindcvtaerj.lpbgm;/z")
	// Layouts["rtna"] = "xdh.qbfoujrtna;gweislkm,/pczyv"
	// //Layouts["funny colemaq"] = "'wgdbmhuyqarstplneiozxcfkjv/.,"
	Layouts["workman"] = NewLayout("Workman", "qdrwbjfup;ashtgyneoizxmcvkl,./")
	// Layouts["workman ct"] = "wcldkjyru/ashtmpneoiqvgfbzx',."
	//Layouts["Colby's Funny"] = "/wgdbmho,qarstflneuizxcpkjv'.y"
	//Layouts["ISRT-AI"] = ",lcmkzfuy.arstgpneio;wvdjbh'qx"
	Layouts["halmak"] = NewLayout("Halmak","wlrbz;qudjshnt,.aeoifmvc/gpxky")
	//Layouts["Balance Twelve but Funny"] = "pclmb'uoyknsrtg,aeihzfwdj/.'-x"
	//Layouts["Dynamica 0.1"] = "lfawqzghu,rnoibysetdjp/m'xckv."
	// Layouts["abc"] = "abcdefghijklmnopqrstuvwxyz,./'"
	Layouts["typehack"] = NewLayout("TypeHack", "jghpfqvou;rsntkyiaelzwmdbc,'.x") 
	// Layouts["qgmlwy"] = "qgmlwyfub;dstnriaeohzxcvjkp,./"
	//Layouts["TNWMLC"] = "tnwmlcbprhsgxjfkqzv;eadioyu,./"
	//Layouts["0.1"] = NewLayout("0.1", "vlafqzgu,ytronbmdeiskj/hpcw'.x")
	//Layouts["0.2"] = NewLayout("0.2", "ydlwkzfuo,strmcbneaiqj'gvph/x.")
	//Layouts["0.2mb"] = NewLayout("0.2mb", "kdl.gxfuoystrm,pneaivz'cwbh/qj")
	//Layouts["0.3"] = NewLayout("0.3", "kfawxqbulytsodchnerizv'gmp.,j/")
	//Layouts["0.4"] = NewLayout("0.4", "ymlkjqfau,scrtdbnoeixw'gvph/z.")
	//Layouts["0.5"] = NewLayout("0.5", "yluwqkfha.sredcmtnoixj'gpzvb/,")
	//Layouts["0.6"] = NewLayout("0.6", ".yuwfqzalvisedcmnort/x,gpbh'jk") // -rolling, +index balance
	//Layouts["0.7"] = NewLayout("0.7", "yhavzqwulfinotkcders/b.mjpg,'x")
	//Layouts["0.7a"] = NewLayout("0.7a", "yauvzqwhlfioetkcdnrs/.,mjpgb'x")
	//Layouts["0.7a3"] = NewLayout("0.7a3", "ylhvzqwuofirntkcdeas/'bmjpg,.x")
	Layouts["rolll"] = NewLayout("rolll", "yauwbxkclvioenpdhsrtj/,.qfmg'z")
	Layouts["semimak"] = NewLayout("Semimak", "flhvzqwuoysrntkcdeaix'bmjpg,./")
	//Layouts["1.0x"] = NewLayout("1.0x", "flhvzqwuoysrntkcdeai'xbmjpg,./")

	// Layouts["0.7mv"] = NewLayout("0.7mv", "yhamzqwulfinotkcders/b.vjpg,'x")
	// Layouts["0.7idk"] = NewLayout("0.7idk", "yhamkqwulfinotvcders/b.jzpg,'x")

	Layouts["whorf"] = NewLayout("Whorf", "flhdmvwou,srntkgyaeixjbzqpc';.")
	Layouts["strtyp"] = NewLayout("strtyp", "jyuozkdlcwhiea,gtnsr'x/.qpbmfv")

	Layouts["flaw"] = NewLayout("FLAW", "flawpzkur/hsoycmtenibj'gvqd.x,")
	// Layouts["beakl"] = "qyouxgcrfzkhea.dstnbj/,i'wmlpv"
	// Layouts["owomak"] = "qwfpbjluy;arstdhneioxvcbzkm,./"
	//Layouts["boo"] = NewLayout("Boo", ",.ucvzfmlyaoesgpntri;x'djbhkwq") // old version, deprecated
	Layouts["boo"] = NewLayout("Boo", ",.ucvzfdlyaoesgpntri;x'wjbhmkq")

	Layouts["x1"] = NewLayout("X1", "kyo,'fclpvhieaudstnrzqj.;wgmbx")
	// Layouts["colemake"] = ";lgwvqpdu.arstkfnhio,jcmzb'y/x"
	// //Layouts["ctgap"] = "qwgdbmhuy'orstplneiazxcfkjv/.,"
	Layouts["ctgap"] = NewLayout("CTGAP", "wcldkjyou/rsthmpneiazvgfbqx',.")
	Layouts["aptap"] = NewLayout("APTAP", "wcdl'/youqrsthmpneiavbgk,.fjxz")
	//Layouts["qwerg"] = NewLayout("QWERG", "qwergkpmb;asdtvyhniozcflxju,./")
	// Layouts["rsthd"] = "jcyfkzl,uqrsthdmnaio/vgpbxw.;-"
	// Layouts["notgate"] = "youwg.vdlpiaescmhtrn'q;xzf,kjb"
	// Layouts["slider"] = "qwfpbjvuyzarscgmneio'ktdxlh/.,"
	// Layouts["paper 200"] = " wldk mic asthy nero bgf vpuj "

	//trigrams := Trigrams(Layouts["mtgap"].Keys)

	if len(args) > 0 {
		if args[0] == "a" || args[0] == "analyze" {
			if len(args) == 1 {
				fmt.Println("You must provide the name of a layout to analyze")
				os.Exit(1)
			}
			Data = LoadData()

			input := strings.ToLower(args[1])
			PrintAnalysis(Layouts[input])
		} else if args[0] == "r" {
			Data = LoadData()

			type x struct {
				name string
				score float64
			}

			var sorted []x

			for _, v := range Layouts {
				sorted = append(sorted, x{v.Name, Score(v)})
			}

			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].score < sorted[j].score
			})

			for _, l := range sorted {
				spaces := strings.Repeat(".", 20-len(l.name))
				fmt.Printf("%s.%s%.2f\n", l.name, spaces, l.score)
			}
		} else if args[0] == "g" {
			Data = LoadData()
			start := time.Now()
			best := Populate(1000)
			end := time.Now()
			fmt.Println(end.Sub(start))

			optimal := Score(best)

			type x struct {
				name string
				score float64
			}

			var sorted []x

			for k, v := range Layouts {
				sorted = append(sorted, x{k, Score(v)})
			}

			sort.Slice(sorted, func(i, j int) bool {
				return sorted[i].score < sorted[j].score
			})
			
			for _, l := range sorted {
				spaces := strings.Repeat(".", 25-len(l.name))
				fmt.Printf("%s.%s%d%%\n", l.name, spaces, int(100*optimal/(Score(Layouts[l.name]))))
			}

		} else if args[0] == "sfbs" {
			Data = LoadData()
			if len(args) == 1 {
				fmt.Println("You must specify a layout")
				os.Exit(1)
			}
			input := strings.ToLower(args[1])
			l := Layouts[input]
			total := 100*float64(SFBs(l.Keys))/float64(Data.TotalBigrams)
			sfbs := ListSFBs(l.Keys)
			SortFreqList(sfbs)
			fmt.Printf("%.2f%%\n", total)
			PrintFreqList(sfbs, 16)
		} else if args[0] == "dsfbs" {
			Data = LoadData()
			if len(args) == 1 {
				fmt.Println("You must specify a layout")
				os.Exit(1)
			}
			input := strings.ToLower(args[1])
			l := Layouts[input]
			total := 100*float64(DSFBs(l.Keys))/float64(Data.TotalBigrams)
			dsfbs := ListDSFBs(l.Keys)
			SortFreqList(dsfbs)
			fmt.Printf("%.2f%%\n", total)
			PrintFreqList(dsfbs, 16)
		}else if args[0] == "bigrams" {
			Data = LoadData()
			if len(args) == 1 {
				fmt.Println("You must specify a layout")
				os.Exit(1)
			}
			input := strings.ToLower(args[1])
			l := Layouts[input]
			sf := ListWeightedSameFinger(l.Keys)
			SortFreqList(sf)
			PrintFreqList(sf, 16)
		} else if args[0] == "dynamic" {
			Data = LoadData()
			if len(args) == 1 {
				fmt.Println("You must specify a layout")
				os.Exit(1)
			}
			input := strings.ToLower(args[1])
			l := Layouts[input]
			truecount, usage := SFBsMinusTop(l.Keys)
			total := 100*float64(usage)/float64(Data.TotalBigrams)
			dynamic, truesfbs := ListRepeats(l.Keys)
			SortFreqList(dynamic)
			SortFreqList(truesfbs)
			fmt.Printf("Dynamic Usage: %.2f%%\n", total)
			PrintFreqList(dynamic, 30)
			fmt.Printf("True SFBs: %.2f%%\n", 100*float64(truecount)/float64(Data.TotalBigrams))
			PrintFreqList(truesfbs, 8)
		} else if args[0] == "speed" {
			Data = LoadData()
			input := strings.ToLower(args[1])
			l := Layouts[input]
			speeds := FingerSpeed(l.Keys)
			fmt.Println("Unweighted Speed")
			for i, v := range speeds {
				fmt.Printf("\t%s: %.2f\n", FingerNames[i], v)
			}
			
		} else if args[0] == "h" {			
			Data = LoadData()
			Heatmap(Layouts[args[1]].Keys)
		} else if args[0] == "ngram" {
			Data = LoadData()
			total := float64(Data.Total)
			ngram := args[1]
			if len(ngram) == 1 {
				fmt.Printf("unigram: %.3f%%\n", 100*float64(Data.Letters[ngram]) / total)
			} else if len(ngram) == 2 {
				fmt.Printf("bigram: %.3f%%\n", 100*float64(Data.Bigrams[ngram]) / total)
				fmt.Printf("skipgram: %.3f%%\n", 100*Data.Skipgrams[ngram] / total)
			} else if len(ngram) == 3 {
				fmt.Printf("trigram: %.3f%%\n", 100*float64(Data.Trigrams[ngram]) / total)
			}
			// } else if args[0] == "i" {
			// 	LoadData()
			// 	input := strings.ToLower(args[1])
			// 	InteractiveAnalysis(Layouts[input])
		} else if args[0] == "load" {
			Data = GetTextData()
			WriteData(Data)
		}
	}
}
