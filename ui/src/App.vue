<template>
  <form action ="">
    <input type="text" >
    <input type="button" value="request">
  </form>
</template>

<script>
import axios from 'axoios'
export default {
  name: 'App',
  getName (){
    axios.get("/objects/objects1/:name",{
      output: 'json'
    })
  }
}

export default {
  name: 'App',
  
}
methods:{
  postreq: function{
          axios({ method: "POST", url: "http://127.0.0.1:8090/calc", data: data, headers: {"content-type": "text/plain" } }).then(result => { })
          // this.response = result.data;
        
}
          }


</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
