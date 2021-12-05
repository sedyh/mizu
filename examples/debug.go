package main

//func main() {
//	_ = ers()
//}
//
//func ers() error {
//	return nil
//}

//img := ebiten.NewImage(16, 16)
//img = img.SubImage(image.Rect(8, 8, 16, 16)).(*ebiten.Image)
//
//got, _ := SubImage(img, image.Rect(5, 5, 4, 4)).(*ebiten.Image).Size()
//want := 4
//fmt.Printf("got: %v, want: %v", got, want)

//imgA := image.NewRGBA(image.Rect(0, 0, 16, 16))
//imgB := imgA.SubImage(image.Rect(8, 8, 16, 16)).(*image.RGBA)
//imgC := SubImage(imgB, image.Rect(1, 1, 4, 4)).(*image.RGBA)

//if err := ebiten.RunGame(engine.NewWorld(&Menu{})); err != nil {
//	log.Fatal(err)
//}

//
//type SubImager interface {
//	SubImage(r image.Rectangle) image.Image
//}
//
//// RelativeSubImage return cropped image relative to the previous crop
//// If you try to make a sub image from another sub image, you may run into
//// unusual behavior due to the fact that absolute coordinates are used for each crop,
//// instead you can use this function so as not to constantly remember this nuance.
//func RelativeSubImage(img image.Image, r image.Rectangle) image.Image {
//	if source, ok := img.(SubImager); ok {
//		rx, ry := img.Bounds().Min.X+r.Min.X, img.Bounds().Min.Y+r.Min.Y
//		return source.SubImage(image.Rect(rx, ry, rx+r.Max.X, ry+r.Max.Y))
//	}
//	return img
//}

//type Game struct {
//	engine.World
//}
//
//func NewGame(first engine.Scene) *Game {
//	return &Game{engine.NewWorld(first)}
//}
//
//func (g *Game) Update() error {
//	return g.Update()
//}
//
//func (g *Game) Draw(screen *ebiten.Image) {
//	g.Draw(screen)
//}
//
//func (g *Game) Layout(w, h int) (int, int) {
//	return g.Layout(w, h)
//}
