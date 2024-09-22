import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";


export const addUser =()=>{

  
    const qury = useQueryClient();


    return  useMutation({
      mutationFn: async (data) => {
        const res = await axios.post("http://localhost:8080/post",data);
  
        return res.data;
       
      },
      onSuccess: async (data) => {
        qury.invalidateQueries({
          queryKey: ["users"],
        });
        console.log("real data", data);
      },
      // onError: async (er) => {
      //   console.log(er.request.response);
      // },
    });
    
}    