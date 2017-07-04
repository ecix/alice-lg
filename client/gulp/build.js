'use strict';

/**
 * Gulp main configuration
 */

var fs    = require('fs');
var path  = require('path');
var gulp  = require('gulp');

var runSequence = require('run-sequence');

var tasks = fs.readdirSync('./gulp/tasks').filter(function(filename){
  return path.extname(filename) === '.js';
});

// == Load configuration
global.config = JSON.parse(fs.readFileSync('./gulp/config.json'));

// == Import all tasks
tasks.forEach(function(task) {
  require('./tasks/' + task);
});


// == Register build task.
gulp.task('build', [
  'bundle',
  'html',
  'stylesheets',
  'assets',
  'app'
]);


// == Production task
gulp.task('production', function(cb) {
  process.env.NODE_ENV = 'production';
  runSequence('default', cb);
});

// == Register default task
gulp.task('default', ['clean'], function(cb) {
  runSequence('build', cb);
});

