module.exports = function(grunt) {

  // ===========================================================================
  // CONFIGURE GRUNT ===========================================================
  // ===========================================================================
  grunt.initConfig({

    // set the varibel for output javascript directory
    distConcatOutput: 'js/bundle.js',
    distMinifiedOutput: 'js/bundle.min.js',
    javascriptTestCase: 'js/src/index.html',
    // banner: "/*! <%= pkg.name %> - v<%= pkg.version %> - ' + '<%= grunt.template.today(\"yyyy-mm-dd\") %> */",
    banner: "/*! '<%= grunt.template.today(\"yyyy-mm-dd\") %> */",

    // set the varibale for input javascript directory
    distInput: 'js/src/*.js',

    // get the configuration info from package.json ----------------------------
    // this way we can use things like name and version (pkg.name)
    // pkg: grunt.file.readJSON('package.json'),


    //configure concat plugins
    concat: {
        options:{
            separator: ';',
            banner: '<%= banner %>'
        },
        dist:{
            src: ['<%= distInput %>'],
            dest: '<%= distConcatOutput %>'
        },
    },
    // configure jshint to validate js files -----------------------------------
    jshint: {
      options: {
        reporter: require('jshint-stylish')
      },
      all: ['Grunfile.js', '<%= distInput %>']
    },

    qunit_junit : {
        options: {
            dest: 'report/',
        }
    },

    qunit: {
		options: {
			"--web-security": "no",
			// coverage: {
			// 	src: [ '<%= distInput %>' ],
			// 	instrumentedFiles: "temp/",
			// 	htmlReport: "report/coverage",
			// 	lcovReport: "report/lcov",
			// 	linesThresholdPct: 10
			// }
		},
		files: ["js/tests/*.html"]
        // all: ["js/tests/*.html"]
    },

    // clean the bundle file
    clean: {
        js: ['<%= distConcatOutput %>', '<%= distMinifiedOutput %>']
    },

    // configure uglify to minify js files -------------------------------------
    uglify: {
      options: {
        banner: '<%= banner %>',
      },
      build: {
        files: {
          '<%= distMinifiedOutput %>': ['<%= distInput %>']
        }
      }
    },

    // configure watch to auto update ------------------------------------------
    watch: {
      scripts: {
        files: '<%= distInput %>',
        tasks: ['uglify', 'concat']
      }
    }

  });

  // ===========================================================================
  // LOAD GRUNT PLUGINS ========================================================
  // ===========================================================================
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks("grunt-coveralls");
  // grunt.loadNpmTasks('grunt-contrib-qunit');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks("grunt-qunit-istanbul");
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-qunit-junit');


  // ===========================================================================
  // CREATE TASKS ==============================================================
  // ===========================================================================
  grunt.registerTask('testqunit', ['qunit_junit', 'qunit']);
  grunt.registerTask('default', ['uglify', 'concat']);
  grunt.registerTask('dev', ['watch']);

};

// https://github.com/npm/npm/issues/10013
// https://github.com/gruntjs/grunt-contrib-uglify
// https://github.com/gruntjs/grunt-contrib-concat

//istanbul issue: https://github.com/gruntjs/grunt-contrib-qunit/issues/139
