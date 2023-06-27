package bigint

func Simplify(num string) string {
	for (string(num[0])=="0"&&len(num)!=1) ||(string(num[0])=="+") {
		num=num[1:]
	}
	for (string(num[0])=="-"&&len(num)!=1&&string(num[0])=="0") {
		num=num[:1]+num[2:]
	}
	return num

}