package helper

// need to hardcode a default template
// given yaegi's constraints on finding files on disk
// provided by this plugin
func GetDefaultCapJsTmpl() string {
	return `<html>
  <head>
    <title>Verifying connection</title>
    <script src="{{ .FrontendJS }}" async defer referrerpolicy="no-referrer"></script>
  </head>
  <body>
    <h1>Verifying connection</h1>
    <p>One moment while we verify your network connection.</p>
    <cap-widget
      id="default"
      data-cap-api-endpoint="{{ .ChallengeURL }}">
      onsolve="console.log(`Token: ${event.detail.token}`)"
    </cap-widget>
  </body>
</html>`
}
