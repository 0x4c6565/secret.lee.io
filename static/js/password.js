var PasswordGenerator = {
	length: 12,
	lowercase: true,
	uppercase: true,
	numbers: true,
	symbols: true,
	setLength: function(length) {
		this.length = length;
	},
	getRandom: function() {
		var result = new Uint32Array(1);
	    window.crypto.getRandomValues(result);
	    return (result[0]/(0xffffffff + 1));
	},
	shuffleArray: function(array) {
		for (var i = array.length - 1; i > 0; i--) {
	        var j = Math.floor(this.getRandom() * (i + 1));
	        var temp = array[i];
	        array[i] = array[j];
	        array[j] = temp;
	    }

	    return array;
	},
	generate: function() {

	    var lowerChars = "abcdefghijklmnopqrstuvwxyz";
	    var upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
		var numberChars = "0123456789";
		var symbolChars = "!$%^&*()_-?+=";

		var allChars = ""
	    var randPasswordArray = Array(this.length);
		var idx = 0;
		function addChars(chars) {
			allChars+=chars;
			randPasswordArray[idx]=chars;
			idx++;
		}

		if (this.lowercase) {
			addChars(lowerChars);
		}

		if (this.uppercase) {
			addChars(upperChars);
		}

		if (this.numbers) {
			addChars(numberChars);
		}

		if (this.symbols) {
			addChars(symbolChars);
		}

	    randPasswordArray = randPasswordArray.fill(allChars, idx);
	    var y=this;
	    return this.shuffleArray(randPasswordArray.map(function(x) { return x[Math.floor(y.getRandom() * x.length)] })).join('');
	}
}
