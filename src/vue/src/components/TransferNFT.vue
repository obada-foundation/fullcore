<template>
  <div>
    <div class="sp-type-form sp-box sp-shadow">
      <form class="sp-voter__main__form">
        <div class="sp-type-form__field sp-form-group">
          <input type="text" class="sp-input" placeholder="DID" v-model="did" />
        </div>

        <div class="sp-type-form__field sp-form-group">
          <input type="text" class="sp-input" placeholder="Receiver address" v-model="receiver" />
        </div>

        <div class="sp-type-form__btns">
          <sp-button @click="transferNFT" :disabled="!isValid" :busy="inFlight">Transfer</sp-button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "TransferNFT",
  data() {
    return {
      inFlight: false,
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
    },
    isValid() {
      return this.did.length > 0 && this.receiver.length > 0
    },
  },
  methods: {
    async transferNFT() {
      this.inFlight = true

      const value = {
        sender: this.currentAccount,
        receiver: this.receiver,
        did: this.did,
      };

      const resp = await this.$store.dispatch("obadafoundation.fullcore.obit/sendMsgSend", {
        value,
        fee: [],
      });

      this.did      = ""
      this.receiver = ""
      this.inFlight = false
    },
  },
}
</script>