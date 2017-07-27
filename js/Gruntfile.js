module.exports = function(grunt) {

  // ===========================================================================
  // CONFIGURE GRUNT ===========================================================
  // ===========================================================================
  grunt.initConfig({

    // set the varibel for output javascript directory
    distConcatOutput: './bundle.js',
    distMinifiedOutput: './bundle.min.js',
    javascriptTestCase: './src/index.html',
    banner: "/*! '<%= grunt.template.today(\"yyyy-mm-dd\") %> */",

    // set the varibale for input javascript directory
    distInput: './src/*.js',

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
      options:{
          jshintrc: './../.jshintrc'
      },
      all: ['Grunfile.js', '<%= distInput %>']
    },

    qunit: {
        options: {
            timeout: 30000,
            "--web-security": "no",
            coverage: {
                src: [ "<%= distInput %>" ],
                instrumentedFiles: "./../temp/",
                htmlReport: "./coverage/html",
                lcovReport: "./coverage/lcov",
                coberturaReport: "./coverage/cobertura",
                linesThresholdPct: 0
            }
        },
        all: ["./tests/*.html"]
    },

    // clean the bundle file
    clean: {
        js: ['<%= distConcatOutput %>', '<%= distMinifiedOutput %>', './coverage/']
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
        tasks: ['jshint', 'uglify', 'concat']
      }
    }

  });

  // ===========================================================================
  // LOAD GRUNT PLUGINS ========================================================
  // ===========================================================================
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-coveralls');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks("grunt-qunit-istanbul");


  // ===========================================================================
  // CREATE TASKS ==============================================================
  // ===========================================================================
  grunt.registerTask('default', ['uglify', 'concat']);
  grunt.registerTask('dev', ['watch']);
};
