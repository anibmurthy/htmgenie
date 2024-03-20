package mapper

var Mapper map[string]func(string) string

func init() {
	//Load all applicable functions handling updates
	Mapper = make(map[string]func(string) string)
	Mapper["#"] = Heading1
	Mapper["##"] = Heading2
	Mapper["###"] = Heading3
	Mapper["####"] = Heading4
	Mapper["#####"] = Heading5
	Mapper["######"] = Heading6
}
