<html>
    <head>
    <title>Kristina Portfolio UI</title>
    <link rel="stylesheet" href="styles/main.css" />
    </head>
    <body>
        <h1>{{ .Name }}</h1>
        {{ range .Projects }}
            <a href="{{ .HTMLURL }}">{{ .HTMLURL }}</a>
            <p>{{ .Description }}</p>
        {{ end }}
        <br />
        <br />
        <img src="http://cultofthepartyparrot.com/parrots/hd/gentlemanparrot.gif" />
    </body>
</html>
