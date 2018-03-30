"use strict";

var webpackConfig = require('./webpack.config.js');

/* global module */
module.exports = function gruntInit(grunt) {

    // ========== GRUNT INIT CONFIG ==========
    grunt.initConfig({
        distHTMLs: './html/*.html', // all html files
        distInput: './js/elems/*.js',   // input source files
        distConcatOutput: './js/bundle.js', // output bundle
        distMinifiedOutput: './js/bundle.min.js',   // output bundle in minified version
        banner: "/*! '<%= grunt.template.today(\"yyyy-mm-dd\") %> */",  // banner for output file

        //configure concat plugins
        concat: {
            options:{
                separator: '\n',
                banner: '<%= banner %>'
            },
            dist:{
                src: ['<%= distInput %>'],
                dest: '<%= distConcatOutput %>'
            }
        },

        // validate files with JSHINT
        jshint: {
            // REF: https://glebbahmutov.com/blog/linting-js-inside-html/
            html: {
                options: {
                    extract: 'always',
                    browser: true,
                },
                files: {
                    src: ['<%= distHTMLs %>']
                }
            },
            js: {
                src: ['Gruntfile.js', '<%= distInput %>']
            },
            options:{
                jshintrc: '.jshintrc'
            },
            // beforeconcat: ['<%= distInput %>'],
            // afterconcat: ['<%= distConcatOutput %>'], // when concatOutput is created then only run lint on it
            // all: ['Grunfile.js', '<%= distInput %>']
        },

        // clean the files, created on temporary based
        clean: {
            js: ['<%= distConcatOutput %>', '<%= distMinifiedOutput %>', './coverage/']
        },

        // minification options
        uglify: {
            options: {
                banner: '<%= banner %>'
            },
            build: {
                files: {
                    '<%= distMinifiedOutput %>': ['<%= distInput %>']
                }
            }
        },

        // run tasks whenever watched files changed
        watch: {
            scripts: {
                files: ['<%= distInput %>', '<%= distHTMLs %>', '.jshintrc'],
                tasks: ['clean', 'jshint', 'concat', 'uglify']
            }
        }

        // Load webpack configuration
        // webpack: {
        //     options: {
        //         stats: !process.env.NODE_ENV || process.env.NODE_ENV === 'development'
        //     },
        //     prod: webpackConfig,
        //     dev: Object.assign({ watch: false }, webpackConfig)
        // }

    }); // initConfig::END

    // ========== AVAILABLE TASKS ==========
    grunt.loadNpmTasks('grunt-contrib-clean');
    grunt.loadNpmTasks('grunt-contrib-jshint');
    grunt.loadNpmTasks('grunt-contrib-concat');
    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-contrib-watch');
    // grunt.loadNpmTasks('grunt-webpack'); // To run webpack via grunt

    // ========== REGISTERED TASKS ==========
    grunt.registerTask('default', ['clean', 'jshint', 'concat', 'uglify']);
    grunt.registerTask('dev', ['watch']);
};
