	<!-- start sidebar -->
	<div id="sidebar">
		<ul>
			<li>
				<h2>{{if .Automezzi}}AUTOMEZZI{{end}}</h2>
				<ul>
					{{if .Automezzi}}{{.AutomezziMenu}}{{end}}
				</ul>
			
			</li>
			
		</ul>
	</div>
	<!-- end sidebar -->