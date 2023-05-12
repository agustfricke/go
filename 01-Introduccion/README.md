Go es un lenguaje de programación compilado y de tipo estático que a menudo 
se describe como el "C del siglo XXI"
Fue diseñado para ser eficiente en términos de tiempo de ejecución, 
fácil de aprender y escribir.

// Go fue creado por expertos 
// como Ken Thompson. el inventor de lengujes como B y C, La 
// idea de crear un nuevo lenguaje de programación surgió en 2007, 
// cuando Google estaba lidiando con problemas de escalabilidad y 
// productividad en su infraestructura de software.

// Los programadores de Google encontraron que los lenguajes de programación existentes no 
// satisfacían completamente sus necesidades, ya que algunos eran demasiado 
// lentos, otros eran demasiado complejos y difíciles de aprender.

// Entonces, el equipo comenzó a trabajar en un nuevo lenguaje de programación que 
// pudiera abordar estas limitaciones y ofrecer una mejor productividad y escalabilidad,
// La versión 1.0 se lanzó como open source en 2012.

// Go se inspira en otros lenguajes como C, C++ y Python, y se ha ganado la 
// popularidad en la comunidad de desarrolladores debido a su simplicidad, eficiencia y capacidad 
// para manejar cargas de trabajo intensivas.

// Go es un lenguaje de programación compilado, 
// lo que significa que el código fuente se traduce a 
// binario, lo que proporciona un 
// mejor rendimiento. También tiene un recolector de basura 
// incorporado que administra automáticamente la memoria, lo que 
// ayuda a evitar problemas comunes de programación, 
// como los errores de asignación y liberación de memoria.

// Go es de tipo estático esto se refiere a la comprobación de tipos en tiempo de compilación, en lugar de 
// en tiempo de ejecución. Esto significa que el compilador de Go verifica que 
// las variables y funciones estén utilizando los tipos correctos antes de que el programa se ejecute.

// En resumen, Go es un lenguaje de programación moderno y poderoso que 
// se utiliza en una amplia variedad de aplicaciones, desde servidores web 
// hasta herramientas de línea de comandos y sistemas operativos.

<head>
	<title>Download and Install Go</title>
</head>
	<h3>Download and install Go</h3>
	<p>Download and install Go quickly with the steps described <a href="#">here</a>.</p>
	<p>For other content on installing, you might be interested in:</p>
	<ul>
		<li>Managing Go installations -- How to install multiple versions and uninstall.</li>
		<li>Installing Go from source -- How to check out the sources, build them on your own machine, and run them.</li>
	</ul>
	<h4>Download (1.20.4)</h4>
	<p>Don't see your operating system here? Try one of the other downloads.</p>
	<h4>Go installation</h4>
	<p>Select the tab for your computer's operating system below, then follow its installation instructions.</p>
	<code>Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:</code>
	<pre>$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz</pre>
	<p>(You may need to run the command as root or through sudo).</p>
	<p>Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.</p>
	<p>Add /usr/local/go/bin to the PATH environment variable.</p>
	<p>You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):</p>
	<code>export PATH=$PATH:/usr/local/go/bin</code>
	<p>Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.</p>
	<p>Verify that you've installed Go by opening a command prompt and typing the following command:</p>
	<code>$ go version</code>
	<p>Confirm that the command prints the installed version of Go.</p>
	<p>You're all set!</p>
	<p>Visit the <a href="#">Getting Started</a> tutorial to write some simple Go code. It takes about 10 minutes to complete.</p>
