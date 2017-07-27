"use strict";
function defineDateFmts() {
    var month = (new Date()).getMonth() + 1;
    var year  = (new Date()).getFullYear();
    // // US Format
    $('input[type=us-date]').w2field('date', { format: 'm/d/yyy' });
    $('input[type=us-dateA]').w2field('date', { format: 'm/d/yyyy', start:  month + '/5/' + year, end: month + '/25/' + year });
    $('input[type=us-dateB]').w2field('date', { format: 'm/d/yyyy', blocked: [ month+'/12/2014',month+'/13/2014',month+'/14/' + year,]});
    $('input[type=us-date1]').w2field('date', { format: 'm/d/yyyy', end: $('input[type=us-date2]') });
    $('input[type=us-date2]').w2field('date', { format: 'm/d/yyyy', start: $('input[type=us-date1]') });
    $('input[type=us-time]').w2field('time',  { format: 'h12' });
    $('input[type=us-timeA]').w2field('time', { format: 'h12', start: '8:00 am', end: '4:30 pm' });

    // EU Common Format
    $('input[type=eu-date]').w2field('date',  { format: 'd.m.yyyy' });
    $('input[type=eu-dateA]').w2field('date', { format: 'd.m.yyyy', start:  '5.' + month + '.' + year, end: '25.' + month + '.' + year });
    $('input[type=eu-dateB]').w2field('date', { format: 'd.m.yyyy', blocked: ['12.' + month + '.' + year, '13.' + month + '.' + year, '14.' + month + '.' + year]});
    $('input[type=eu-date1]').w2field('date', { format: 'd.m.yyyy', end: $('input[type=eu-date2]') });
    $('input[type=eu-date2]').w2field('date', { format: 'd.m.yyyy', start: $('input[type=eu-date1]') });
    $('input[type=eu-time]').w2field('time',  { format: 'h24' });
    $('input[type=eu-timeA]').w2field('time', { format: 'h24', start: '8:00 am', end: '4:30 pm' });
}
