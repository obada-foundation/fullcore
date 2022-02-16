<template>
  <div class="sp-type-form sp-box sp-shadow">
    <form class="sp-type-form__main__form">
      <div class="sp-type-form__header sp-box-header">Create an NFT</div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="serial_number_hash" placeholder="Serial Number" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="manufacturer" placeholder="Manufacturer" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="part_number" placeholder="Part Number" required />
      </div>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="owner_did" placeholder="Physical Asset Owner ID" />
      </div>
      <SpButton type="primary" v-on:click="getToken">Get Token</SpButton>

      <div class="sp-type-form__field sp-form-group">
        <input type="text" class="sp-input" v-model="trust_anchor" placeholder="Trust Anchor (Registered Agent) ID" disabled />
      </div>

      <div class="sp-type-form__btns">
        <SpButton type="primary" v-on:click="submit">Mint OBT</SpButton>
      </div>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "MintNftForm",
  data() {
    return {
      // Until we not define trust-anchor spec with Denis and Joe we going to use sandbox service
      taAuthToken: "eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiZWU0MjRmN2QtODVjNi00MTM2LWJmMjQtYTgxYjA3ZGI2OTY5In0.3sl5tmp_BOmHSRpQbWMSK5CpDmHQLNOHpy1WzNgrkvDbEi4P3DdpiD-1iH9QsQajx_JZhkRq3S4-Ti1Cs390Bw",
      trustAnchors: [{
        id: "demoTrustAnchor",
        name: "Demo Trust Anchor"
      }],
      serial_number_hash: "",
      manufacturer: "",
      part_number: "",
      owner_did: "",
      trust_anchor: "demoTrustAnchor",
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
    getToken() {
      axios.post("http://demo.ta.alpha.obada.io/api/v1/issue-token", {}, {
        headers: {
          'Authrization': 'Bearer ' + this.taAuthToken,
          'Content-Type': 'application/json'
        },
        withCredentials: true
      })
          .then(response => console.log(response))
          .catch(error => {
            this.errorMessage = error.message;
            console.error("There was an error!", error);
          });

      return ""
    },
    async submit() {
      const value = {
        creator: this.currentAccount,
        serialNumberHash: this.serial_number_hash,
        manufacturer: this.manufacturer,
        partNumber: this.part_number,
        ownerDid: this.owner_did,
        trustAnchor: this.trust_anchor,
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
