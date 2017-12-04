<html>
    <head>
    <title>Kristina Portfolio UI</title>
    <link rel="stylesheet" href="styles/main.css" />
    </head>
    <body>
        <h1>{{ .Name }}</h1>
        <h2>{{ .Heading }}</h2>
        <img src="http://krisvincent.com/img/kristina.jpg" alt="ProfilePic">
        {{ range .Projects }}            
            <div>{{ .Name }}:</div>
            <ul>
            <li><a href="{{ .HTMLURL }}">{{ .HTMLURL }}</a></li>
            <li>Description: {{ .Description }}</li>
            <li>Written in: {{ .Language }}</li>
            </ul>
        {{ end }}
        <br />
        <br />
    </body>
</html>
