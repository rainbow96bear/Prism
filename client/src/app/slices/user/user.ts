import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";

interface User {
  id: string;
}

const initialState: User = {
  id: "",
};

// 비동기 작업을 처리하는 thunk 생성
export const getUserInfo = createAsyncThunk<User>(
  "userInfo/getProfile",
  async () => {
    const response = await axios.get<User>(`/users/info`, {
      withCredentials: true,
    });
    return response.data; // API 응답에서 필요한 데이터를 반환
  }
);

export const logout = createAsyncThunk("auth/logout", async () => {
  await axios.post(
    "/oauth/kakao/logout",
    {},
    {
      withCredentials: true,
    }
  );

  // 로그아웃 후 리턴되는 데이터를 사용하려면 여기에 추가 코드 작성 가능
});

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getUserInfo.fulfilled, (state, action) => {
        // 비동기 작업이 성공하면 상태를 업데이트
        state.id = action.payload.id;
      })
      .addCase(getUserInfo.rejected, (state, action) => {
        state.id = "";
        console.error("Error fetching user info:", action.error.message);
      })
      .addCase(logout.fulfilled, (state) => {
        // 로그아웃이 성공하면 상태를 초기화
        state.id = "";
      })
      .addCase(logout.rejected, (state, action) => {
        console.error("Logout failed:", action.error.message);
      });
  },
});

export default userSlice.reducer;
