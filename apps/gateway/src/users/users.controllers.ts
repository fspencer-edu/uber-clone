import { Body, Controller, Get, Param, Post } from '@nestjs/common';
import axios from 'axios';

const USER_SERVICE_URL =
  process.env.USER_SERVICE_URL || 'http://user-service:8080';

@Controller('users')
export class UsersController {
  @Get()
  async getUsers() {
    const response = await axios.get(`${USER_SERVICE_URL}/users`);
    return response.data;
  }

  @Get(':id')
  async getUserById(@Param('id') id: string) {
    const response = await axios.get(`${USER_SERVICE_URL}/users/${id}`);
    return response.data;
  }

  @Post()
  async createUser(
    @Body() body: { name: string; email: string }
  ) {
    const response = await axios.post(`${USER_SERVICE_URL}/users`, body);
    return response.data;
  }
}