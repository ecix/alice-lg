'use strict';

/**
 * Task: app
 *
 * Compile the main app
 */

var gulp       = require('gulp');
var gulpif     = require('gulp-if');
var uglify     = require('gulp-uglify');
var rename     = require('gulp-rename');
var size       = require('gulp-size');

var browserify = require('browserify');
var babelify   = require('babelify');

var source     = require('vinyl-source-stream');
var buffer     = require('vinyl-buffer');

// == Register task: app
gulp.task('app', function() {
  var production = process.env.NODE_ENV == 'production';
  var entries = ['./app.jsx'];

  var bundler = browserify({
    debug: !production,
    entries: entries,
    extensions: ['.jsx'],
    paths: ['./node_modules', './']
  });

  bundler.transform('envify', { global: true });
  bundler.transform(babelify.configure({
    presets: ["es2015", "react"]
  }));


  return bundler.bundle()
    .pipe(source('app.js'))
    .pipe(gulpif(production, buffer()))
    .pipe(gulpif(production, uglify()))
    .pipe(gulp.dest('build/js'))
    .pipe(size());
});
