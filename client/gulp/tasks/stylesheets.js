'use strict';

/**
 * Task: stylesheets 
 *
 * Compile stylesheets from less source.
 */

var gulp         = require('gulp');
var gulpif       = require('gulp-if');
var sass         = require('gulp-sass');
var cssmin       = require('gulp-cssmin');
var rename       = require('gulp-rename');
var autoprefixer = require('gulp-autoprefixer');



// == Register task: stylesheets 
gulp.task('stylesheets', function(){
  var production = process.env.NODE_ENV == 'production';

  // Compile less files
  return gulp.src('assets/scss/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(autoprefixer({
      browsers: ['last 2 versions'],
      cascade: false
     }))
    .pipe(gulpif(production, cssmin()))
    .pipe(gulp.dest('build/css/'));
});
