import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "../../../configs/AxiosConfig";
import { Admin } from "../../../GlobalType/Admin";

interface AdminState {
  isAdmin: boolean;
  admin_info: Admin;
  done: boolean;
}

const initialState: AdminState = {
  isAdmin: false,
  admin_info: { id: "", rank: 0 },
  done: false,
};
// 비동기 작업을 처리하는 thunk 생성
export const getAdminAuth = createAsyncThunk<AdminState>(
  "admin/getAdminAuth",
  async () => {
    const response = await axios.get<AdminState>(`/admins/authorization`, {
      withCredentials: true,
    });
    return response.data; // API 응답에서 필요한 데이터를 반환
  }
);

export const login = createAsyncThunk<Admin, string>(
  "admin/login",
  async (password) => {
    const response = await axios.post<Admin>(
      `/admins/login`,
      { password: password },
      {
        withCredentials: true,
      }
    );
    return response.data; // API 응답에서 필요한 데이터를 반환
  }
);

export const logout = createAsyncThunk<Admin>("admin/logout", async () => {
  const response = await axios.post<Admin>(
    `/admins/logout`,
    {},
    {
      withCredentials: true,
    }
  );
  return response.data; // API 응답에서 필요한 데이터를 반환
});

const adminSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getAdminAuth.fulfilled, (state, action) => {
        state.isAdmin = action.payload.isAdmin;
        state.admin_info.id = action.payload.admin_info.id;
        state.admin_info.rank = action.payload.admin_info.rank;
        state.done = true;
      })
      .addCase(getAdminAuth.rejected, (state, action) => {
        state.isAdmin = false;
        state.admin_info.id = "";
        state.admin_info.rank = 0;
        state.done = false;
        console.error("Error fetching user info:", action.error.message);
      })
      .addCase(login.fulfilled, (state, action) => {
        state.admin_info.id = action.payload.id;
        state.admin_info.rank = action.payload.rank;
        state.done = true;
      })
      .addCase(login.rejected, (state, action) => {
        state.admin_info.id = "";
        state.admin_info.rank = 0;
        state.done = false;
        console.error("Error fetching user info:", action.error.message);
      })
      .addCase(logout.fulfilled, (state, action) => {
        state.admin_info.id = "";
        state.admin_info.rank = 0;
        state.done = true;
      })
      .addCase(logout.rejected, (state, action) => {
        state.admin_info.id = "";
        state.admin_info.rank = 0;
        state.done = false;
        console.error("Error fetching user info:", action.error.message);
      });
  },
});

export default adminSlice.reducer;
