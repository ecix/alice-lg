'use strict';

/**
 * Task: bundle
 *
 * Bundle dependencies into a single file.
 *
 * See: config.bundle
 */

var gulp        = require('gulp');
var gulpif      = require('gulp-if');
var rename      = require('gulp-rename');
var concat      = require('gulp-concat');
var cssmin      = require('gulp-cssmin');
var uglify      = require('gulp-uglify');
var mergeStream = require('merge-stream');


gulp.task('bundle', ['bundle-js', 'bundle-css']);

gulp.task('bundle-js', function() {
  var production = process.env.NODE_ENV == 'production';
  var js = global.config.bundle.js;
  var stream = mergeStream();

  if(js) {
    for(var name in js) {
      if(js.hasOwnProperty(name)) {
        var files = js[name];
        stream.add(
          gulp.src(files)
            .pipe(concat(name + '.js'))
            .pipe(gulpif(production, uglify()))
            .pipe(gulp.dest('build/js')));
      }
    }
  }

  return stream.isEmpty() ? null : stream;
});

gulp.task('bundle-css', function() {
  var css = global.config.bundle.css;
  var production = process.env.NODE_ENV == 'production';
  var stream = mergeStream();

  if(css) {
    for(var name in css) {
      if(css.hasOwnProperty(name)) {
        var files = css[name];
        stream.add(
          gulp.src(files)
            .pipe(concat(name + '.css'))
            .pipe(gulpif(production, cssmin()))
            .pipe(gulp.dest('build/css')));
      }
    }
  }

  return stream.isEmpty() ? null : stream;
});
