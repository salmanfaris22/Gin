
import { useQuery } from "@tanstack/react-query";
import axios from "axios";






export const useUserList=()=>{

    return useQuery({
        queryKey: ["users"],
        queryFn: async () => {
          const res = await axios.get("http://localhost:8080/posts");
          return res.data;
        },
      })
    
    }
  
