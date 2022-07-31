package utils

func Ensure(err error){
	if err!=nil{
		panic(err)
	}
}
