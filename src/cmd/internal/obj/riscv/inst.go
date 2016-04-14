// Automatically generated by parse-opcodes

package riscv

import "cmd/internal/obj"

type inst struct {
	opcode uint32
	funct3 uint32
	csr int64
	funct7 uint32
}

func encode(a obj.As) (i *inst, ok bool) {
	switch a {
	case ABEQ:
		return &inst{0x63, 0x0, 0, 0x0}, true
	case ABNE:
		return &inst{0x63, 0x1, 0, 0x0}, true
	case ABLT:
		return &inst{0x63, 0x4, 0, 0x0}, true
	case ABGE:
		return &inst{0x63, 0x5, 0, 0x0}, true
	case ABLTU:
		return &inst{0x63, 0x6, 0, 0x0}, true
	case ABGEU:
		return &inst{0x63, 0x7, 0, 0x0}, true
	case AJALR:
		return &inst{0x67, 0x0, 0, 0x0}, true
	case AJAL:
		return &inst{0x6f, 0x0, 0, 0x0}, true
	case ALUI:
		return &inst{0x37, 0x0, 0, 0x0}, true
	case AAUIPC:
		return &inst{0x17, 0x0, 0, 0x0}, true
	case AADDI:
		return &inst{0x13, 0x0, 0, 0x0}, true
	case ASLLI:
		return &inst{0x13, 0x1, 0, 0x0}, true
	case ASLTI:
		return &inst{0x13, 0x2, 0, 0x0}, true
	case ASLTIU:
		return &inst{0x13, 0x3, 0, 0x0}, true
	case AXORI:
		return &inst{0x13, 0x4, 0, 0x0}, true
	case ASRLI:
		return &inst{0x13, 0x5, 0, 0x0}, true
	case ASRAI:
		return &inst{0x13, 0x5, 1024, 0x20}, true
	case AORI:
		return &inst{0x13, 0x6, 0, 0x0}, true
	case AANDI:
		return &inst{0x13, 0x7, 0, 0x0}, true
	case AADD:
		return &inst{0x33, 0x0, 0, 0x0}, true
	case ASUB:
		return &inst{0x33, 0x0, 1024, 0x20}, true
	case ASLL:
		return &inst{0x33, 0x1, 0, 0x0}, true
	case ASLT:
		return &inst{0x33, 0x2, 0, 0x0}, true
	case ASLTU:
		return &inst{0x33, 0x3, 0, 0x0}, true
	case AXOR:
		return &inst{0x33, 0x4, 0, 0x0}, true
	case ASRL:
		return &inst{0x33, 0x5, 0, 0x0}, true
	case ASRA:
		return &inst{0x33, 0x5, 1024, 0x20}, true
	case AOR:
		return &inst{0x33, 0x6, 0, 0x0}, true
	case AAND:
		return &inst{0x33, 0x7, 0, 0x0}, true
	case AADDIW:
		return &inst{0x1b, 0x0, 0, 0x0}, true
	case ASLLIW:
		return &inst{0x1b, 0x1, 0, 0x0}, true
	case ASRLIW:
		return &inst{0x1b, 0x5, 0, 0x0}, true
	case ASRAIW:
		return &inst{0x1b, 0x5, 1024, 0x20}, true
	case AADDW:
		return &inst{0x3b, 0x0, 0, 0x0}, true
	case ASUBW:
		return &inst{0x3b, 0x0, 1024, 0x20}, true
	case ASLLW:
		return &inst{0x3b, 0x1, 0, 0x0}, true
	case ASRLW:
		return &inst{0x3b, 0x5, 0, 0x0}, true
	case ASRAW:
		return &inst{0x3b, 0x5, 1024, 0x20}, true
	case ALB:
		return &inst{0x3, 0x0, 0, 0x0}, true
	case ALH:
		return &inst{0x3, 0x1, 0, 0x0}, true
	case ALW:
		return &inst{0x3, 0x2, 0, 0x0}, true
	case ALD:
		return &inst{0x3, 0x3, 0, 0x0}, true
	case ALBU:
		return &inst{0x3, 0x4, 0, 0x0}, true
	case ALHU:
		return &inst{0x3, 0x5, 0, 0x0}, true
	case ALWU:
		return &inst{0x3, 0x6, 0, 0x0}, true
	case ASB:
		return &inst{0x23, 0x0, 0, 0x0}, true
	case ASH:
		return &inst{0x23, 0x1, 0, 0x0}, true
	case ASW:
		return &inst{0x23, 0x2, 0, 0x0}, true
	case ASD:
		return &inst{0x23, 0x3, 0, 0x0}, true
	case AFENCE:
		return &inst{0xf, 0x0, 0, 0x0}, true
	case AFENCEI:
		return &inst{0xf, 0x1, 0, 0x0}, true
	case AMUL:
		return &inst{0x33, 0x0, 32, 0x1}, true
	case AMULH:
		return &inst{0x33, 0x1, 32, 0x1}, true
	case AMULHSU:
		return &inst{0x33, 0x2, 32, 0x1}, true
	case AMULHU:
		return &inst{0x33, 0x3, 32, 0x1}, true
	case ADIV:
		return &inst{0x33, 0x4, 32, 0x1}, true
	case ADIVU:
		return &inst{0x33, 0x5, 32, 0x1}, true
	case AREM:
		return &inst{0x33, 0x6, 32, 0x1}, true
	case AREMU:
		return &inst{0x33, 0x7, 32, 0x1}, true
	case AMULW:
		return &inst{0x3b, 0x0, 32, 0x1}, true
	case ADIVW:
		return &inst{0x3b, 0x4, 32, 0x1}, true
	case ADIVUW:
		return &inst{0x3b, 0x5, 32, 0x1}, true
	case AREMW:
		return &inst{0x3b, 0x6, 32, 0x1}, true
	case AREMUW:
		return &inst{0x3b, 0x7, 32, 0x1}, true
	case AAMOADDW:
		return &inst{0x2f, 0x2, 0, 0x0}, true
	case AAMOXORW:
		return &inst{0x2f, 0x2, 512, 0x10}, true
	case AAMOORW:
		return &inst{0x2f, 0x2, 1024, 0x20}, true
	case AAMOANDW:
		return &inst{0x2f, 0x2, 1536, 0x30}, true
	case AAMOMINW:
		return &inst{0x2f, 0x2, -2048, 0x40}, true
	case AAMOMAXW:
		return &inst{0x2f, 0x2, -1536, 0x50}, true
	case AAMOMINUW:
		return &inst{0x2f, 0x2, -1024, 0x60}, true
	case AAMOMAXUW:
		return &inst{0x2f, 0x2, -512, 0x70}, true
	case AAMOSWAPW:
		return &inst{0x2f, 0x2, 128, 0x4}, true
	case ALRW:
		return &inst{0x2f, 0x2, 256, 0x8}, true
	case ASCW:
		return &inst{0x2f, 0x2, 384, 0xc}, true
	case AAMOADDD:
		return &inst{0x2f, 0x3, 0, 0x0}, true
	case AAMOXORD:
		return &inst{0x2f, 0x3, 512, 0x10}, true
	case AAMOORD:
		return &inst{0x2f, 0x3, 1024, 0x20}, true
	case AAMOANDD:
		return &inst{0x2f, 0x3, 1536, 0x30}, true
	case AAMOMIND:
		return &inst{0x2f, 0x3, -2048, 0x40}, true
	case AAMOMAXD:
		return &inst{0x2f, 0x3, -1536, 0x50}, true
	case AAMOMINUD:
		return &inst{0x2f, 0x3, -1024, 0x60}, true
	case AAMOMAXUD:
		return &inst{0x2f, 0x3, -512, 0x70}, true
	case AAMOSWAPD:
		return &inst{0x2f, 0x3, 128, 0x4}, true
	case ALRD:
		return &inst{0x2f, 0x3, 256, 0x8}, true
	case ASCD:
		return &inst{0x2f, 0x3, 384, 0xc}, true
	case ASCALL:
		return &inst{0x73, 0x0, 0, 0x0}, true
	case ASBREAK:
		return &inst{0x73, 0x0, 1, 0x0}, true
	case ASRET:
		return &inst{0x73, 0x0, 256, 0x8}, true
	case ASFENCEVM:
		return &inst{0x73, 0x0, 257, 0x8}, true
	case AWFI:
		return &inst{0x73, 0x0, 258, 0x8}, true
	case AMRTH:
		return &inst{0x73, 0x0, 774, 0x18}, true
	case AMRTS:
		return &inst{0x73, 0x0, 773, 0x18}, true
	case AHRTS:
		return &inst{0x73, 0x0, 517, 0x10}, true
	case ACSRRW:
		return &inst{0x73, 0x1, 0, 0x0}, true
	case ACSRRS:
		return &inst{0x73, 0x2, 0, 0x0}, true
	case ACSRRC:
		return &inst{0x73, 0x3, 0, 0x0}, true
	case ACSRRWI:
		return &inst{0x73, 0x5, 0, 0x0}, true
	case ACSRRSI:
		return &inst{0x73, 0x6, 0, 0x0}, true
	case ACSRRCI:
		return &inst{0x73, 0x7, 0, 0x0}, true
	case AFADDS:
		return &inst{0x53, 0x0, 0, 0x0}, true
	case AFSUBS:
		return &inst{0x53, 0x0, 128, 0x4}, true
	case AFMULS:
		return &inst{0x53, 0x0, 256, 0x8}, true
	case AFDIVS:
		return &inst{0x53, 0x0, 384, 0xc}, true
	case AFSGNJS:
		return &inst{0x53, 0x0, 512, 0x10}, true
	case AFSGNJNS:
		return &inst{0x53, 0x1, 512, 0x10}, true
	case AFSGNJXS:
		return &inst{0x53, 0x2, 512, 0x10}, true
	case AFMINS:
		return &inst{0x53, 0x0, 640, 0x14}, true
	case AFMAXS:
		return &inst{0x53, 0x1, 640, 0x14}, true
	case AFSQRTS:
		return &inst{0x53, 0x0, 1408, 0x2c}, true
	case AFADDD:
		return &inst{0x53, 0x0, 32, 0x1}, true
	case AFSUBD:
		return &inst{0x53, 0x0, 160, 0x5}, true
	case AFMULD:
		return &inst{0x53, 0x0, 288, 0x9}, true
	case AFDIVD:
		return &inst{0x53, 0x0, 416, 0xd}, true
	case AFSGNJD:
		return &inst{0x53, 0x0, 544, 0x11}, true
	case AFSGNJND:
		return &inst{0x53, 0x1, 544, 0x11}, true
	case AFSGNJXD:
		return &inst{0x53, 0x2, 544, 0x11}, true
	case AFMIND:
		return &inst{0x53, 0x0, 672, 0x15}, true
	case AFMAXD:
		return &inst{0x53, 0x1, 672, 0x15}, true
	case AFCVTSD:
		return &inst{0x53, 0x0, 1025, 0x20}, true
	case AFCVTDS:
		return &inst{0x53, 0x0, 1056, 0x21}, true
	case AFSQRTD:
		return &inst{0x53, 0x0, 1440, 0x2d}, true
	case AFLES:
		return &inst{0x53, 0x0, -1536, 0x50}, true
	case AFLTS:
		return &inst{0x53, 0x1, -1536, 0x50}, true
	case AFEQS:
		return &inst{0x53, 0x2, -1536, 0x50}, true
	case AFLED:
		return &inst{0x53, 0x0, -1504, 0x51}, true
	case AFLTD:
		return &inst{0x53, 0x1, -1504, 0x51}, true
	case AFEQD:
		return &inst{0x53, 0x2, -1504, 0x51}, true
	case AFCVTWS:
		return &inst{0x53, 0x0, -1024, 0x60}, true
	case AFCVTWUS:
		return &inst{0x53, 0x0, -1023, 0x60}, true
	case AFCVTLS:
		return &inst{0x53, 0x0, -1022, 0x60}, true
	case AFCVTLUS:
		return &inst{0x53, 0x0, -1021, 0x60}, true
	case AFMVXS:
		return &inst{0x53, 0x0, -512, 0x70}, true
	case AFCLASSS:
		return &inst{0x53, 0x1, -512, 0x70}, true
	case AFCVTWD:
		return &inst{0x53, 0x0, -992, 0x61}, true
	case AFCVTWUD:
		return &inst{0x53, 0x0, -991, 0x61}, true
	case AFCVTLD:
		return &inst{0x53, 0x0, -990, 0x61}, true
	case AFCVTLUD:
		return &inst{0x53, 0x0, -989, 0x61}, true
	case AFMVXD:
		return &inst{0x53, 0x0, -480, 0x71}, true
	case AFCLASSD:
		return &inst{0x53, 0x1, -480, 0x71}, true
	case AFCVTSW:
		return &inst{0x53, 0x0, -768, 0x68}, true
	case AFCVTSWU:
		return &inst{0x53, 0x0, -767, 0x68}, true
	case AFCVTSL:
		return &inst{0x53, 0x0, -766, 0x68}, true
	case AFCVTSLU:
		return &inst{0x53, 0x0, -765, 0x68}, true
	case AFMVSX:
		return &inst{0x53, 0x0, -256, 0x78}, true
	case AFCVTDW:
		return &inst{0x53, 0x0, -736, 0x69}, true
	case AFCVTDWU:
		return &inst{0x53, 0x0, -735, 0x69}, true
	case AFCVTDL:
		return &inst{0x53, 0x0, -734, 0x69}, true
	case AFCVTDLU:
		return &inst{0x53, 0x0, -733, 0x69}, true
	case AFMVDX:
		return &inst{0x53, 0x0, -224, 0x79}, true
	case AFLW:
		return &inst{0x7, 0x2, 0, 0x0}, true
	case AFLD:
		return &inst{0x7, 0x3, 0, 0x0}, true
	case AFSW:
		return &inst{0x27, 0x2, 0, 0x0}, true
	case AFSD:
		return &inst{0x27, 0x3, 0, 0x0}, true
	case AFMADDS:
		return &inst{0x43, 0x0, 0, 0x0}, true
	case AFMSUBS:
		return &inst{0x47, 0x0, 0, 0x0}, true
	case AFNMSUBS:
		return &inst{0x4b, 0x0, 0, 0x0}, true
	case AFNMADDS:
		return &inst{0x4f, 0x0, 0, 0x0}, true
	case AFMADDD:
		return &inst{0x43, 0x0, 32, 0x1}, true
	case AFMSUBD:
		return &inst{0x47, 0x0, 32, 0x1}, true
	case AFNMSUBD:
		return &inst{0x4b, 0x0, 32, 0x1}, true
	case AFNMADDD:
		return &inst{0x4f, 0x0, 32, 0x1}, true
	case ASLLIRV32:
		return &inst{0x13, 0x1, 0, 0x0}, true
	case ASRLIRV32:
		return &inst{0x13, 0x5, 0, 0x0}, true
	case ASRAIRV32:
		return &inst{0x13, 0x5, 1024, 0x20}, true
	case AFRFLAGS:
		return &inst{0x73, 0x2, 1, 0x0}, true
	case AFSFLAGS:
		return &inst{0x73, 0x1, 1, 0x0}, true
	case AFSFLAGSI:
		return &inst{0x73, 0x5, 1, 0x0}, true
	case AFRRM:
		return &inst{0x73, 0x2, 2, 0x0}, true
	case AFSRM:
		return &inst{0x73, 0x1, 2, 0x0}, true
	case AFSRMI:
		return &inst{0x73, 0x5, 2, 0x0}, true
	case AFSCSR:
		return &inst{0x73, 0x1, 3, 0x0}, true
	case AFRCSR:
		return &inst{0x73, 0x2, 3, 0x0}, true
	case ARDCYCLE:
		return &inst{0x73, 0x2, -1024, 0x60}, true
	case ARDTIME:
		return &inst{0x73, 0x2, -1023, 0x60}, true
	case ARDINSTRET:
		return &inst{0x73, 0x2, -1022, 0x60}, true
	case ARDCYCLEH:
		return &inst{0x73, 0x2, -896, 0x64}, true
	case ARDTIMEH:
		return &inst{0x73, 0x2, -895, 0x64}, true
	case ARDINSTRETH:
		return &inst{0x73, 0x2, -894, 0x64}, true
	case AECALL:
		return &inst{0x73, 0x0, 0, 0x0}, true
	case AEBREAK:
		return &inst{0x73, 0x0, 1, 0x0}, true
	case AERET:
		return &inst{0x73, 0x0, 256, 0x8}, true
	}
	return nil, false
}
