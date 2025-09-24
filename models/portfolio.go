package models

type Contact struct {
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
}

type CTA struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type Profile struct {
	ProfilePic string  `json:"profile_pic"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
	Intro      string  `json:"intro"`
	Location   string  `json:"location"`
	Contact    Contact `json:"contact"`
	CTA        CTA     `json:"cta"`
}

type Experience struct {
	Company     string `json:"company"`
	Role        string `json:"role"`
	Description string `json:"description"`
	StartYear   string `json:"start_year"`
	EndYear     string `json:"end_year"`
	Link        string `json:"link"`
}

type Project struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	Year         string   `json:"year"`
	Link         string   `json:"link"`
}

type Skills struct {
	DevopsAndCICD       []string `json:"devops_and_ci/cd"`
	AIAndGenerativeAI   []string `json:"ai_and_generative_ai"`
	ProgrammingLanguages []string `json:"programming_languages"`
	Backend             []string `json:"backend"`
}

type Education struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Year        string `json:"year"`
	Link        string `json:"link"`
}

type Social struct {
	Platform string `json:"platform"`
	Link     string `json:"link"`
	Icon     string `json:"icon"`
}

type PortfolioData struct {
	Profile    Profile      `json:"profile"`
	Experience []Experience `json:"experience"`
	Projects   []Project    `json:"projects"`
	Skills     Skills       `json:"skills"`
	Education  []Education  `json:"education"`
	Socials    []Social     `json:"socials"`
}