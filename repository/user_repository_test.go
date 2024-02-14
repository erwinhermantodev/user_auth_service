package repository

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_CreateUser(t *testing.T) {
	
	tests := []struct {
		name          string
		inputUser     *User
		mockQuery     string
		mockArgs      []interface{}
		mockRows      []interface{}
		expectedID    int
		expectedError error
	}{
		{
			name:       "Success",
			inputUser:  &User{PhoneNumber: "123456789", FullName: "John Doe", Password: "password123"},
			mockQuery:  "INSERT INTO users",
			mockArgs:   []interface{}{"123456789", "John Doe", sqlmock.AnyArg()},
			mockRows:   []interface{}{1},
			expectedID: 1,
			
			expectedError: nil,
		},
		
	}

	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			
			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error initializing mock database: %v", err)
			}
			defer mockDB.Close()

			
			ur := NewUserRepository(mockDB)

			args := make([]driver.Value, len(tc.mockArgs))
			for i, v := range tc.mockArgs {
				args[i] = v
			}

			rowValues := make([]driver.Value, len(tc.mockRows))
			for i, v := range tc.mockRows {
				
				value, ok := v.(int)
				if !ok {
					t.Fatalf("Unexpected type in mockRows. Expected int, got %T", v)
				}
				rowValues[i] = int64(value)
			}


			
			mock.ExpectQuery(tc.mockQuery).
				WithArgs(args...).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(rowValues...))
			
			userID, err := ur.CreateUser(tc.inputUser)

			
			if userID != tc.expectedID {
				t.Errorf("Unexpected user ID. Expected: %d, Got: %d", tc.expectedID, userID)
			}
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Unexpected error. Expected: %v, Got: %v", tc.expectedError, err)
			}

			
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled mock expectations: %v", err)
			}
		})
	}
}

func TestUserRepository_FindUserByPhone(t *testing.T) {
    
    tests := []struct {
        name           string
        mockQuery      string
        mockArgs       []interface{}
        mockRows       []interface{}
        expectedUser   *User
        expectedError  error
    }{
        {
            name:         "User found",
            mockQuery:    "SELECT id, phone_number, full_name, password FROM users WHERE phone_number",
            mockArgs:     []interface{}{"123456789"},
            mockRows:     []interface{}{1, "123456789", "John Doe", "hashed_password"},
            expectedUser: &User{ID: 1, PhoneNumber: "123456789", FullName: "John Doe", Password: "hashed_password"},
            expectedError: nil,
        },
        {
            name:         "User not found",
            mockQuery:    "SELECT id, phone_number, full_name, password FROM users WHERE phone_number",
            mockArgs:     []interface{}{"987654321"},
            mockRows:     []interface{}{1, "123456789", "John Doe", "hashed_password"},
            expectedUser: &User{ID: 1, PhoneNumber: "123456789", FullName: "John Doe", Password: "hashed_password"},
            expectedError: nil, 
        },
    }

    
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            
            mockDB, mock, err := sqlmock.New()
            if err != nil {
                t.Fatalf("Error initializing mock database: %v", err)
            }
            defer mockDB.Close()

            
            ur := NewUserRepository(mockDB)

            
            args := make([]driver.Value, len(tc.mockArgs))
            for i, v := range tc.mockArgs {
                args[i] = v.(driver.Value)
            }

            
            rowValues := make([]driver.Value, len(tc.mockRows))
            for i, v := range tc.mockRows {
                switch val := v.(type) {
                case int:
                    rowValues[i] = int64(val) 
                case string:
                    rowValues[i] = val
                
                default:
                    t.Fatalf("Unexpected type in mockRows: %T", v)
                }
            }

            
            rows := sqlmock.NewRows([]string{"id", "phone_number", "full_name", "password"})
            rows.AddRow(rowValues...)
            mock.ExpectQuery(tc.mockQuery).
                WithArgs(args...).
                WillReturnRows(rows)

            
            user, err := ur.FindUserByPhone(tc.mockArgs[0].(string))

            
            if (user == nil && tc.expectedUser != nil) || (user != nil && tc.expectedUser == nil) || (user != nil && tc.expectedUser != nil && *user != *tc.expectedUser) {
                t.Errorf("Unexpected user. Expected: %v, Got: %v", tc.expectedUser, user)
            }
            if !errors.Is(err, tc.expectedError) {
                t.Errorf("Unexpected error. Expected: %v, Got: %v", tc.expectedError, err)
            }

            
            if err := mock.ExpectationsWereMet(); err != nil {
                t.Errorf("Unfulfilled mock expectations: %v", err)
            }
        })
    }
}


func TestUserRepository_GetUserByID(t *testing.T) {
    
    tests := []struct {
        name           string
        userID         int
        mockQuery      string
        mockArgs       []interface{}
        mockRows       []interface{}
        expectedUser   *User
        expectedError  error
    }{
        {
            name:          "User found",
            userID:        1,
            mockQuery:     "SELECT id, phone_number, full_name, password FROM users WHERE id",
            mockArgs:      []interface{}{1},
            mockRows:      []interface{}{1, "123456789", "John Doe", "hashed_password"},
            expectedUser:  &User{ID: 1, PhoneNumber: "123456789", FullName: "John Doe", Password: "hashed_password"},
            expectedError: nil,
        },
        {
            name:          "User not found",
            userID:        2,
            mockQuery:     "SELECT id, phone_number, full_name, password FROM users WHERE id",
            mockArgs:      []interface{}{2},
            mockRows:      []interface{}{1, "123456789", "John Doe", "hashed_password"},
            expectedUser:  &User{ID: 1, PhoneNumber: "123456789", FullName: "John Doe", Password: "hashed_password"},
            expectedError: nil, 
        },
        
    }

    
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            
            mockDB, mock, err := sqlmock.New()
            if err != nil {
                t.Fatalf("Error initializing mock database: %v", err)
            }
            defer mockDB.Close()

            
            ur := NewUserRepository(mockDB)
			args := make([]driver.Value, len(tc.mockArgs))
            for i, v := range tc.mockArgs {
                args[i] = v.(driver.Value)
            }

			rowValues := make([]driver.Value, len(tc.mockRows))
            for i, v := range tc.mockRows {
                switch val := v.(type) {
                case int:
                    rowValues[i] = int64(val) 
                case string:
                    rowValues[i] = val
                
                default:
                    t.Fatalf("Unexpected type in mockRows: %T", v)
                }
            }

            
            mock.ExpectQuery(tc.mockQuery).
                WithArgs(args...).
                WillReturnRows(sqlmock.NewRows([]string{"id", "phone_number", "full_name", "password"}).AddRow(rowValues...))

            
            user, err := ur.GetUserByID(tc.userID)

            
            if (user == nil && tc.expectedUser != nil) || (user != nil && tc.expectedUser == nil) || (user != nil && tc.expectedUser != nil && *user != *tc.expectedUser) {
                t.Errorf("Unexpected user. Expected: %v, Got: %v", tc.expectedUser, user)
            }
            if !errors.Is(err, tc.expectedError) {
                t.Errorf("Unexpected error. Expected: %v, Got: %v", tc.expectedError, err)
            }

            
            if err := mock.ExpectationsWereMet(); err != nil {
                t.Errorf("Unfulfilled mock expectations: %v", err)
            }
        })
    }
}

func TestUserRepository_UpdateUserProfile(t *testing.T) {
    
    tests := []struct {
        name          string
        inputUser     *User
        mockQuery     string
        mockArgs      []interface{}
        expectedError error
    }{
        {
            name:          "Success",
            inputUser:     &User{ID: 1, FullName: "Jane Doe"},
            mockQuery:     "UPDATE users SET full_name",
            mockArgs:      []interface{}{"Jane Doe", 1},
            expectedError: nil,
        },
        
    }

    
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            
            mockDB, mock, err := sqlmock.New()
            if err != nil {
                t.Fatalf("Error initializing mock database: %v", err)
            }
            defer mockDB.Close()

            
            ur := NewUserRepository(mockDB)
			args := make([]driver.Value, len(tc.mockArgs))
            for i, v := range tc.mockArgs {
                args[i] = v.(driver.Value)
            }

            
            mock.ExpectExec(tc.mockQuery).
                WithArgs(args...).
                WillReturnResult(sqlmock.NewResult(0, 1)) 

            
            err = ur.UpdateUserProfile(tc.inputUser)

            
            if !errors.Is(err, tc.expectedError) {
                t.Errorf("Unexpected error. Expected: %v, Got: %v", tc.expectedError, err)
            }

            
            if err := mock.ExpectationsWereMet(); err != nil {
                t.Errorf("Unfulfilled mock expectations: %v", err)
            }
        })
    }
}