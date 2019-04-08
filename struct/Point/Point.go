package Point

type Point struct {
    View int16
    Like int16
    Stock int16
    Comment int16
}

func GetInitPoint() *Point {
    p := new(Point)

    p.View = 1
    p.Like = 10
    p.Stock = 10
    p.Comment = 10

    return p
}
