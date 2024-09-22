

import { addUser } from "../../hooks/addUSer";
import { useUserList } from "../../hooks/useLogine";


const SignUp = () => {
    const {mutate}=addUser()
   const {data,isLoading,error} =  useUserList()

  

console.log("s",data);

if(isLoading){
  return  <div>Loading....</div>
}
if(error){
    return  <div>Loading....</div>
  }
  return (
    <div>
      <button onClick={()=>mutate( {
          
          title: "jasim muhammed",
          desc: "is Aloso A Boy",
        })}>adduser</button>
    </div>
  )
}


export default SignUp
