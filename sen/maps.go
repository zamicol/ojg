// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen

const (
	skipChar     = 'a'
	skipNewline  = 'b'
	valSlash     = 'c'
	valNeg       = 'f'
	val0         = 'g'
	valDigit     = 'h'
	valQuote     = 'i'
	tokenStart   = 'j'
	openArray    = 'k'
	openObject   = 'l'
	closeArray   = 'm'
	closeObject  = 'n'
	colonColon   = 'q'
	numSpc       = 'r'
	numNewline   = 's'
	numDot       = 't'
	tokenOk      = 'u'
	numFrac      = 'v'
	fracE        = 'w'
	expSign      = 'x'
	expDigit     = 'y'
	strQuote     = 'z'
	negDigit     = '-'
	strSlash     = 'A'
	escOk        = 'B'
	uOk          = 'E'
	tokenSpc     = 'G'
	tokenColon   = 'I'
	tokenNlColon = 'J'
	numDigit     = 'N'
	numZero      = 'O'
	strOk        = 'R'
	escU         = 'U'
	commentStart = 'K'
	commentEnd   = 'L'
	charErr      = '.'

	//   0123456789abcdef0123456789abcdef
	valueMap = "" +
		".........ab..a.................." + // 0x00
		"a.i.j.....jjafjcghhhhhhhhh..j.jj" + // 0x20
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjk.mjj" + // 0x40
		".jjjjjjjjjjjjjjjjjjjjjjjjjjl.nj." + // 0x60
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0x80
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xa0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xc0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjv" //  0xe0
	//   0123456789abcdef0123456789abcdef
	tokenMap = "" +
		".........GJ..G.................." + // 0x00
		"G...u.....uuGuucuuuuuuuuuuI.u.uu" + // 0x20
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuk.muu" + // 0x40
		".uuuuuuuuuuuuuuuuuuuuuuuuuul.nu." + // 0x60
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0x80
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xa0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xc0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuut" //  0xe0
	//   0123456789abcdef0123456789abcdef
	htmlValueMap = "" +
		".........ab..a.................." + // 0x00
		"a.i.j.....jjafjcghhhhhhhhh.....j" + // 0x20
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjk.mjj" + // 0x40
		".jjjjjjjjjjjjjjjjjjjjjjjjjjl.nj." + // 0x60
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0x80
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xa0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj" + // 0xc0
		"jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjv" //  0xe0
	//   0123456789abcdef0123456789abcdef
	htmlTokenMap = "" +
		".........GJ..G.................." + // 0x00
		"G...u.....uuGuucuuuuuuuuuuI....u" + // 0x20
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuk.muu" + // 0x40
		".uuuuuuuuuuuuuuuuuuuuuuuuuul.nu." + // 0x60
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0x80
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xa0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" + // 0xc0
		"uuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuut" //  0xe0
	//   0123456789abcdef0123456789abcdef
	colonMap = "" +
		".........ab..a.................." + // 0x00
		"a.........................q....." + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	negMap = "" +
		"................................" + // 0x00
		"................O---------......" + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	zeroMap = "" +
		".........rs..r.................." + // 0x00
		"r...........r.tc................" + // 0x20
		"...........................k.m.." + // 0x40
		"...........................l.n.." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................n" //  0xe0
	//   0123456789abcdef0123456789abcdef
	digitMap = "" +
		".........rs..r.................." + // 0x00
		"r...........r.tcNNNNNNNNNN......" + // 0x20
		".....w.....................k.m.." + // 0x40
		".....w.....................l.n.." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................n" //  0xe0
	//   0123456789abcdef0123456789abcdef
	dotMap = "" +
		"................................" + // 0x00
		"................vvvvvvvvvv......" + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	fracMap = "" +
		".........rs..r.................." + // 0x00
		"r...........r..cvvvvvvvvvv......" + // 0x20
		".....w.....................k.m.." + // 0x40
		".....w.....................l.n.." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................n" //  0xe0
	//   0123456789abcdef0123456789abcdef
	expSignMap = "" +
		"................................" + // 0x00
		"...........x.x..yyyyyyyyyy......" + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	expZeroMap = "" +
		"................................" + // 0x00
		"................yyyyyyyyyy......" + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	expMap = "" +
		".........rs..r.................." + // 0x00
		"r...........r..cyyyyyyyyyy......" + // 0x20
		"...........................k.m.." + // 0x40
		"...........................l.n.." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................n" //  0xe0
	//   0123456789abcdef0123456789abcdef
	stringMap = "" +
		".........RR..R.................." + // 0x00
		"RRzRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" + // 0x20
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRARRR" + // 0x40
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" + // 0x60
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" + // 0x80
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" + // 0xR0
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" + // 0xc0
		"RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRR" //   0xe0
	//   0123456789abcdef0123456789abcdef
	escMap = "" +
		"................................" + // 0x00
		"..B............B................" + // 0x20
		"............................B..." + // 0x40
		"..B...B.......B...B.BU.........." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	escByteMap = "" +
		"................................" + // 0x00
		"..\"............/................" + // 0x20
		"............................\\..." + // 0x40
		"..\b...\f.......\n...\r.\t.........." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	uMap = "" +
		"................................" + // 0x00
		"................EEEEEEEEEE......" + // 0x20
		".EEEEEE........................." + // 0x40
		".EEEEEE........................." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	spaceMap = "" +
		".........ab..a.................." + // 0x00
		"a...........a..................." + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................s" //  0xe0
	//   0123456789abcdef0123456789abcdef
	commentStartMap = "" +
		"................................" + // 0x00
		"...............K................" + // 0x20
		"................................" + // 0x40
		"................................" + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
	//   0123456789abcdef0123456789abcdef
	commentMap = "" +
		"..........L....................." + // 0x00
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0x20
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0x40
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0x60
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0x80
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0xa0
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" + // 0xc0
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" //   0xe0)
)
