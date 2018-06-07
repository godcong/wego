package message

/*SendLocationInfo SendLocationInfo */
type SendLocationInfo struct {
	LocationX float64 `xml:"LocationX"`
	LocationY float64 `xml:"LocationY"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}
