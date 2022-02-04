<template>
  <div>
    <div class="sp-voter__main sp-box sp-shadow sp-form-group">
        <form class="sp-voter__main__form">
          <div class="sp-voter__main__rcpt__header sp-box-header">
            Create a NFT
          </div>

          <input class="sp-input" placeholder="Serial Number" v-model="serial_number_hash" required />
	  <br><br>

	  <input class="sp-input" placeholder="Manufacturer" v-model="manufacturer" required />
	  <br><br>

	  <input class="sp-input" placeholder="Part Number" v-model="part_number" required />
	  <br><br>

          <input class="sp-input" name="owner_did" placeholder="Physical Asset Owner ID"/>
	  <br /><br>

	  <input class="sp-input" name="trust_anchor" placeholder="Trust Anchor (Registered Agent) ID"  />
	  <br /><br>
 

          <sp-button @click="submit">Mint OBT</sp-button>
        </form>
    </div>
  </div>
</template>
<script>
export default {
  name: "MintNftForm",
  data() {
    return {
      serial_number_hash: "",
      manufacturer: "",
      part_number: "",
      owner_did: "",
      trust_anchor: "",
    };
  },
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
    }
  },
  methods: {
    async submit() {
      const value = {
        creator: this.currentAccount,
        title: this.serial_number_hash,
        manufacturer: this.manufacturer,
        part_number: this.part_number,
        owner_did: this.owner_id,
        trust_anchor: this.trust_anchor,
      };
	
      const resp = await this.$store.dispatch("obadafoundation.fullcore.obit/sendMsgMintObit", {
        value,
        fee: [],
      });

      console.log(resp);
    },
  },
};
</script>
