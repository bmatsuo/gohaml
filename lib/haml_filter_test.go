package gohaml

import "testing"

var filterOutputTests = []io {
	io{":cdata blah\n", "<![CDATA[\n\tblah\n]]>", ""},
	io{":cdata blah\n\tbleh\n", "<![CDATA[\n\tblah\n\tbleh\n]]>", ""},
	io{":cdata blah\n\n\tbleh\n", "<![CDATA[\n\tblah\n\tbleh\n]]>", ""},
	io{":cdata %blah\n", "<![CDATA[\n\t%blah\n]]>", ""},
	io{":cdata %blah\n\t#bleh\n", "<![CDATA[\n\t%blah\n\t#bleh\n]]>", ""},
	io{":cdata\n\tblah\n", "<![CDATA[\n\tblah\n]]>", ""},
	io{":cdata\n\tblah\n\tbleh\n", "<![CDATA[\n\tblah\n\tbleh\n]]>", ""},
	io{":cdata\n\tblah\nbleh\n", "<![CDATA[\n\tblah\n]]>\nbleh", ""},
	io{":cdata\nblah\n", "<![CDATA[\n]]>\nblah", ""},
	io{"%p\n\t:cdata blah\n\t\tbleh\n", "<p>\n\t<![CDATA[\n\t\tblah\n\t\tbleh\n\t]]>\n</p>", ""},
	io{":javascript blah\n", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>", ""},
	io{":javascript blah\n\tbleh\n", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t\tbleh\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah\n", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah\n\tbleh\n", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t\tbleh\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah\nbleh\n", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>\nbleh", ""},
	io{"%p\n\t:javascript blah\n\t\tbleh\n", "<p>\n\t<script type=\"text/javascript\">\n\t\t//<![CDATA[\n\t\t\tblah\n\t\t\tbleh\n\t\t//]]>\n\t</script>\n</p>", ""},
	io{":css blah\n", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>", ""},
	io{":css blah\n\tbleh\n", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t\tbleh\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah\n", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah\n\tbleh\n", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t\tbleh\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah\nbleh\n", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>\nbleh", ""},
	io{"%p\n\t:css blah\n\t\tbleh\n", "<p>\n\t<style type=\"text/css\">\n\t\t/*<![CDATA[*/\n\t\t\tblah\n\t\t\tbleh\n\t\t/*]]>*/\n\t</style>\n</p>", ""},
	// Same tests with no trailing '\n' to make sure both cases work
	io{":cdata blah", "<![CDATA[\n\tblah\n]]>", ""},
	io{":cdata blah\n\tbleh", "<![CDATA[\n\tblah\n\tbleh\n]]>", ""},
	io{":cdata %blah", "<![CDATA[\n\t%blah\n]]>", ""},
	io{":cdata %blah\n\t#bleh", "<![CDATA[\n\t%blah\n\t#bleh\n]]>", ""},
	io{":cdata\n\tblah", "<![CDATA[\n\tblah\n]]>", ""},
	io{":cdata\n\tblah\n\tbleh", "<![CDATA[\n\tblah\n\tbleh\n]]>", ""},
	io{":cdata\n\tblah\nbleh", "<![CDATA[\n\tblah\n]]>\nbleh", ""},
	io{":cdata\nblah", "<![CDATA[\n]]>\nblah", ""},
	io{"%p\n\t:cdata blah\n\t\tbleh", "<p>\n\t<![CDATA[\n\t\tblah\n\t\tbleh\n\t]]>\n</p>", ""},
	io{":javascript blah", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>", ""},
	io{":javascript blah\n\tbleh", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t\tbleh\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah\n\tbleh", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t\tbleh\n\t//]]>\n</script>", ""},
	io{":javascript\n\tblah\nbleh", "<script type=\"text/javascript\">\n\t//<![CDATA[\n\t\tblah\n\t//]]>\n</script>\nbleh", ""},
	io{"%p\n\t:javascript blah\n\t\tbleh", "<p>\n\t<script type=\"text/javascript\">\n\t\t//<![CDATA[\n\t\t\tblah\n\t\t\tbleh\n\t\t//]]>\n\t</script>\n</p>", ""},
	io{":css blah", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>", ""},
	io{":css blah\n\tbleh", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t\tbleh\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah\n\tbleh", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t\tbleh\n\t/*]]>*/\n</style>", ""},
	io{":css\n\tblah\nbleh", "<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t\tblah\n\t/*]]>*/\n</style>\nbleh", ""},
	io{"%p\n\t:css blah\n\t\tbleh", "<p>\n\t<style type=\"text/css\">\n\t\t/*<![CDATA[*/\n\t\t\tblah\n\t\t\tbleh\n\t\t/*]]>*/\n\t</style>\n</p>", ""},
	// Test sequences of filters
	io{":cdata blah\n\tbleh\n:javascript\n:cdata\n\tblah\n:css",
		"<![CDATA[\n\tblah\n\tbleh\n]]>\n<script type=\"text/javascript\">\n\t//<![CDATA[\n\t//]]>\n</script>\n<![CDATA[\n\tblah\n]]>\n<style type=\"text/css\">\n\t/*<![CDATA[*/\n\t/*]]>*/\n</style>", ""},
}

func TestFilterOutput(t *testing.T) {
	for i, io := range filterOutputTests {
		scope := make(map[string]interface{})

		engine, _ := NewEngine(io.input, "\t")
		engine.Autoclose = false
		output := engine.Render(scope)
		if output != io.expected {
			t.Errorf("Test %d:\n\tInput    %q\n\texpected %q\n\tgot      %q", i, io.input, io.expected, output)
			return
		}
	}
}