package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetConsumedLicenses a breakdown of the licenses consumed by an enterprise.
type GetConsumedLicenses struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_seats_consumed property
    total_seats_consumed *int32
    // The total_seats_purchased property
    total_seats_purchased *int32
    // The users property
    users []GetConsumedLicenses_usersable
}
// NewGetConsumedLicenses instantiates a new GetConsumedLicenses and sets the default values.
func NewGetConsumedLicenses()(*GetConsumedLicenses) {
    m := &GetConsumedLicenses{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGetConsumedLicensesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetConsumedLicensesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetConsumedLicenses(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GetConsumedLicenses) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GetConsumedLicenses) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["total_seats_consumed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSeatsConsumed(val)
        }
        return nil
    }
    res["total_seats_purchased"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSeatsPurchased(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGetConsumedLicenses_usersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GetConsumedLicenses_usersable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GetConsumedLicenses_usersable)
                }
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetTotalSeatsConsumed gets the total_seats_consumed property value. The total_seats_consumed property
// returns a *int32 when successful
func (m *GetConsumedLicenses) GetTotalSeatsConsumed()(*int32) {
    return m.total_seats_consumed
}
// GetTotalSeatsPurchased gets the total_seats_purchased property value. The total_seats_purchased property
// returns a *int32 when successful
func (m *GetConsumedLicenses) GetTotalSeatsPurchased()(*int32) {
    return m.total_seats_purchased
}
// GetUsers gets the users property value. The users property
// returns a []GetConsumedLicenses_usersable when successful
func (m *GetConsumedLicenses) GetUsers()([]GetConsumedLicenses_usersable) {
    return m.users
}
// Serialize serializes information the current object
func (m *GetConsumedLicenses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_seats_consumed", m.GetTotalSeatsConsumed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total_seats_purchased", m.GetTotalSeatsPurchased())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetConsumedLicenses) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalSeatsConsumed sets the total_seats_consumed property value. The total_seats_consumed property
func (m *GetConsumedLicenses) SetTotalSeatsConsumed(value *int32)() {
    m.total_seats_consumed = value
}
// SetTotalSeatsPurchased sets the total_seats_purchased property value. The total_seats_purchased property
func (m *GetConsumedLicenses) SetTotalSeatsPurchased(value *int32)() {
    m.total_seats_purchased = value
}
// SetUsers sets the users property value. The users property
func (m *GetConsumedLicenses) SetUsers(value []GetConsumedLicenses_usersable)() {
    m.users = value
}
type GetConsumedLicensesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalSeatsConsumed()(*int32)
    GetTotalSeatsPurchased()(*int32)
    GetUsers()([]GetConsumedLicenses_usersable)
    SetTotalSeatsConsumed(value *int32)()
    SetTotalSeatsPurchased(value *int32)()
    SetUsers(value []GetConsumedLicenses_usersable)()
}
