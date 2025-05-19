package template

const UserHTMTL = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
		body {
			margin: 0;
			font-family: Arial, sans-serif;
			background: #f5f5f5;
			color: #333;
			display: flex;
			flex-direction: column;
			align-items: center;
			padding: 2rem;
		}
		.profile {
			text-align: center;
		}
		.profile img {
			border-radius: 50%;
			width: 100px;
			height: 100px;
			object-fit: cover;
			margin-bottom: 1rem;
		}
		h1 {
			margin: 0.5rem 0 1rem;
		}
		.links {
			width: 100%;
			max-width: 400px;
			display: flex;
			flex-direction: column;
			gap: 1rem;
		}
		.link {
			background: white;
			padding: 1rem;
			border-radius: 8px;
			text-align: center;
			text-decoration: none;
			color: #333;
			font-weight: bold;
			box-shadow: 0 2px 6px rgba(0,0,0,0.1);
			transition: background 0.2s;
		}
		.link:hover {
			background: #eaeaea;
		}
	</style>
</head>
<body>
	<div class="profile">
		<img src="{{.ProfileImageURL}}" alt="Profile Image">
		<h1>{{.UserID}}</h1>
	</div>
	<div class="links">
		{{range .Links}}
			<a class="link" href="{{.URL}}" target="_blank">{{.Label}}</a>
		{{end}}
	</div>
</body>
</html>
`

type Link struct {
	Label string
	URL   string
}

type PageData struct {
	UserID          string
	Title           string
	Password        string
	ProfileImageURL string
	Links           []Link
}
