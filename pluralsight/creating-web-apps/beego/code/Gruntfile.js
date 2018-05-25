module.exports = function(grunt) {

  grunt.initConfig({
    sass: {
      dist: {
        files: {
          'public/css/app.css': 'sass/app.scss'
        }
      }
    },

    watch: {
      scripts: {
        files: ['sass/**/*.scss'],
        tasks: ['sass']
      }
    }
  });

  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

};
