<template>
  <div>
    <h3>List of NFTs</h3>
  </div>
</template>
<style>
  .option-radio > .button {
      height: 40px;
      width: 50%;
  }
</style>
<script>
export default {
    name: 'NftList',
    computed: {
	currentAccount() {
	    if (this._depsLoaded) {
		if (this.loggedIn) {
		    return this.$store.getters['common/wallet/address']
		} else {
		    return null
		}
	    } else {
		return null
	    }
	},
	loggedIn() {
	    if (this._depsLoaded) {
		return this.$store.getters['common/wallet/loggedIn']
	    } else {
		return false
	    }
	},
    },
}
</script>
