   // DOM
   const btn =  document.querySelector('#click')
   const body = document.querySelector('body') 
   const btn_change = document.querySelector('#click1')

   function toggle_Block() {
    body.classList.toggle("open")
   }
   function backgroundcolor() {
      body.style.backgroundColor = "black";  
      
   }
  btn.addEventListener('click', toggle_Block)
  btn_change.addEventListener('click', backgroundcolor)