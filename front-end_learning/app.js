 let N_apo = true;

 N_apo === true ? console.log(`Welcome Student status: ${N_apo}`):console.log("You're not a student ..")

 // Arrow function
 let able = true;
 const Helpmeout = (person) => {
    able === true ? console.log(`yo ${person} i can help you out`): console.log("nah bro i can't help")
 }

 let array = [20,15,13,16];

 let filtrearray = array.filter((number) =>{
    return number >= 18; // checks if the number is sup or equal to 18, if it is it adds the number into a new empty variable 
  }
);
console.log(filtrearray)

let age_array = [18,20,16,15,21];
let agefiltrearray = age_array.filter((age) => {
    return age >= 18;
});
console.log(agefiltrearray)
// without the .filter
for(let i = 0; i < age_array.length-1;i++) {
    if (age_array[i] >= 18) {
        console.log(`${age_array[i]}`);
    }
}
let Us = [10,20,30,40,50]
    let i = 0; 
   while(i != Us.length-1){
        let aud = Us[i]*1.5
        i++
        console.log(aud)
   } 

   // Objects {fancy name for structures i guess :/}
   let users = [
    {
        email: "amaraali@gmail.com",
        password: "nigachain2.o",
        name: "ali",
        discord: "alilacream",
        subscription: "VIP+",
        lessonscompleteds:[true,false]

    },
   ]
   function Create_user(email,password,name,discord,subscription,lessonscompleted) {
            users.push({
                email,
                password,
                name,
                discord,
                subscription,
                lessonscompleted
            });
                return users
   }
   console.log(Create_user("ali@gmail.com","meh","ali","alilacream","Notvip",[false,false], 1))
   // DOM

   document.querySelector('#box1').style.backgroundColor = "red"
   function toggle_Block() {
    document.querySelector('body').classList.toggle("open")
   }