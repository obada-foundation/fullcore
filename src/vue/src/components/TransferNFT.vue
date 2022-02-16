<template>
  <div>
    <div class="sp-voter__main sp-box sp-shadow sp-form-group">
        <form class="sp-voter__main__form">
          <input class="sp-input" placeholder="DID" v-model="did" />
          <br /><br>

          <input class="sp-input" placeholder="Receiver address" v-model="receiver" />
          <br /><br>

          <sp-button @click="submit">Transfer</sp-button>
        </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "TransferNFT",
  data() {
      return {
          did: "",
          receiver: "",
      }
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
        sender: this.currentAccount,
        receiver: this.receiver,
        id: this.did,
      };
	
      const resp = await this.$store.dispatch("obadafoundation.fullcore.obit/sendMsgSend", {
        value,
        fee: [],
      });

      console.log(resp);
    },
  }
}
</script>