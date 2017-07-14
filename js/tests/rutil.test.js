this.jsdom = require('./../../node_modules/jsdom-global')()
global.$ = global.jQuery = require('./../jquery.min.js');

var chai = require('./../../node_modules/chai');
var expect = chai.expect;
var assert = chai.assert;
var rutils = require('./../src/rutil.js');

describe('rutil javascript unit testcase', function(){
    it('plural accord', function(){
        assert.equal(rutils.plural("accord"), "accords");
    });

    it('plural a', function(){
        assert.equal(rutils.plural("a"), "as");
    });

    it('int_to_bool true', function(done){
        assert.equal(rutils.int_to_bool(3), true);
        done();
    });

    it('int_to_bool false', function(){
        assert.equal(rutils.int_to_bool(-1), false);
    });

    it('dateTodayStr success', function(){
        var today = new Date();
        var dd = today.getDate();
        var mm = today.getMonth() + 1; //January is 0!
        var yyyy = today.getFullYear();
        var today_date =  mm + '/' + dd + '/' + yyyy;
        assert.equal(rutils.dateTodayStr(), today_date);
    });

    it('dateTodayStr error', function(){
        expect(rutils.dateTodayStr()).to.not.equal("7/12/2017");
    });

    it('number_format 1234.56', function(){
        assert.equal(rutils.number_format(1234.56), "1,235");
    });

    it('number_format 1234.56-2-', function(){
        assert.equal(rutils.number_format(1234.56, 2, ',', ' '), "1 234,56");
    });

    it('number_format 1234.5678-2-.', function(){
        assert.equal(rutils.number_format(1234.5678, 2, '.', ''), "1234.57");
    });

    it('number_format 67-2-,-.', function(){
        assert.equal(rutils.number_format(67, 2, ',', '.'), "67,00");
    });

    it('number_format 1000', function(){
        assert.equal(rutils.number_format(1000), "1,000");
    });

    it('number_format 67.311, 2', function(){
        assert.equal(rutils.number_format(67.311, 2), "67.31");
    });

    it('number_format 1000.55, 1', function(){
        assert.equal(rutils.number_format(1000.55, 1), "1,000.6");
    });

    it('number_format 67000', function(){
        assert.equal(rutils.number_format(67000, 5, ',', '.'), "67.000,00000");
    });

    it('number_format 0.9, 0', function(){
        assert.equal(rutils.number_format(0.9, 0), "1");
    });

    it('number_format 1.20, 4', function(){
        assert.equal(rutils.number_format('1.20', 4), "1.2000");
    });

    it('number_format 1.2000, 3', function(){
        assert.equal(rutils.number_format('1.2000', 3), "1.200");
    });
});
