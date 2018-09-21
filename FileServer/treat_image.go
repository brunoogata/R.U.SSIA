package main

import(
	"image"
	"image/jpeg"
	"os"
	"log"
	"math"
)

var pixels_base []pixel

type pixel struct {
	r, g, b, a uint32
}

func init(){
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func LoadImage() (image.Image, image.Image){
	imgcore, err := os.Open("./nucleo.jpg")
	imgdata, err2 := os.Open("./teste.jpg")

	if (err != nil) || (err2 != nil) {
		log.Fatal(err)
	}

	base, err_ := jpeg.Decode(imgcore)
	data, _ := jpeg.Decode(imgdata)

	if err_ != nil {
		log.Fatal(err_)
	}

	return base, data
}

func GetArrayPixelsRGBA(pic image.Image, x_bound int, y_bound int) []pixel{
	total_bound := x_bound * y_bound
	pixarr := make([]pixel, total_bound)

	for i := 0; i < total_bound; i++ {
		x := i % x_bound
		y := i / x_bound

		r, g, b, a := pic.At(x,y).RGBA()
		
		pixarr[i].r = r
		pixarr[i].g = g
		pixarr[i].b = b
		pixarr[i].a = a
	}

	return pixarr 
}

func CalculateMeanDistance(pixarr1 []pixel, pixarr2 []pixel, x_bound int, y_bound int) float64{
	deviation := float64(0)
	total_bound := x_bound * y_bound

	for i := 0; i < total_bound; i++ {
		p1_r, p1_g, p1_b, p1_a := pixarr1[i].GetRGBAValues()
		p2_r, p2_g, p2_b, p2_a := pixarr2[i].GetRGBAValues()
		dist := EuclidianDistance(p1_r, p1_g, p1_b, p1_a, p2_r, p2_g, p2_b, p2_a)
		deviation += dist
	}

	return (deviation / float64(total_bound))
}

func SquareDiff(x, y uint32) uint64 {   
    d := uint64(x) - uint64(y)
    return d * d
}

func EuclidianDistance(a1 uint32,b1 uint32,c1 uint32,d1 uint32,a2 uint32,b2 uint32,c2 uint32,d2 uint32) float64{
	dist := SquareDiff(a1,a2) + SquareDiff(b1,b2) + SquareDiff(c1,c2) + SquareDiff(d1,d2)
	dist64 := float64(dist)

	return math.Sqrt(dist64)
}

func (p pixel) GetRGBAValues() (uint32,uint32,uint32,uint32){
	return p.r, p.g, p.b, p.a 
}

func GetStateQueue(base image.Image, data image.Image) string{
	bounds := data.Bounds()

	if(base != nil){
		pixels_base = GetArrayPixelsRGBA(base, bounds.Dx(), bounds.Dy())
	}
	pixels_data := GetArrayPixelsRGBA(data, bounds.Dx(), bounds.Dy())	

	mean_dist := CalculateMeanDistance(pixels_base, pixels_data, bounds.Dx(), bounds.Dy())

	if mean_dist > float64(1000) {
		return "cheio"
	} else {
		return "vazio"
	}
}
