// Listing 11.9  A variety of annotations

type Name struct {
	First string `json:"firstName" xml:"FirstName"`
	Last  string `json:"lastName,omitempty"`
	Other string `not,even.a=tag`
}
