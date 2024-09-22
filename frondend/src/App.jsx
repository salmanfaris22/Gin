// import axios from "axios";
import SignUp from "./components/auth/SignUp";
// import { useMutation, useQuery, useQueryClient } from "react-query";

const App = () => {
  // const qury = useQueryClient();


 
  return (
    <div>
      {/* <button onClick={() => mutate()}>mutate</button> */}
      <h1>Data:</h1>
      <SignUp />
    </div>
  );
};

export default App;
