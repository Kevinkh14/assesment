<template>
  <div class="hello">
    <form v-on:submit.prevent = 'getObject'>
      <input type="text" v-model='input' >
      <Button>enter</Button>
      <h1>Name: {{object.Name}}</h1>
      <h1>Description: {{object.description}}</h1>
      <h1>Attributes: {{object.Attributes}}</h1>
    </form>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'FrontEnd',

  data () {return{
    input :[],
    object:{},
  }},

  methods :{
    getObject() {
      axios.get(`http://localhost:8081/${this.input}`)
       .then(res=>{
         this.object = res.data
         console.log(res.data)
         console.log(this.object.Name)
       })
       .catch(err =>{
         console.log(err)
       })
    }
  }
}
</script>



<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
