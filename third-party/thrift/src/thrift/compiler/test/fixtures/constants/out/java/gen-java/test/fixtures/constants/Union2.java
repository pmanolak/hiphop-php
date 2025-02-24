/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */

package test.fixtures.constants;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import java.util.*;
import org.apache.thrift.*;
import org.apache.thrift.transport.*;
import org.apache.thrift.protocol.*;

import static com.google.common.base.MoreObjects.toStringHelper;

@SwiftGenerated
@ThriftUnion("union2")
public final class Union2 implements com.facebook.thrift.payload.ThriftSerializable {
    
    private static final boolean allowNullFieldValues =
        System.getProperty("thrift.union.allow-null-field-values", "false").equalsIgnoreCase("true");

    private static final TStruct STRUCT_DESC = new TStruct("union2");
    private static final Map<String, Integer> NAMES_TO_IDS = new HashMap();
    public static final Map<String, Integer> THRIFT_NAMES_TO_IDS = new HashMap();
    private static final Map<Integer, TField> FIELD_METADATA = new HashMap<>();
    private static final Union2 _DEFAULT = new Union2();

    public static final int _I = 1;
    private static final TField I_FIELD_DESC = new TField("i", TType.I32, (short)1);
    public static final int _D = 2;
    private static final TField D_FIELD_DESC = new TField("d", TType.DOUBLE, (short)2);
    public static final int _S = 3;
    private static final TField S_FIELD_DESC = new TField("s", TType.STRUCT, (short)3);
    public static final int _U = 4;
    private static final TField U_FIELD_DESC = new TField("u", TType.STRUCT, (short)4);

    static {
      NAMES_TO_IDS.put("i", 1);
      THRIFT_NAMES_TO_IDS.put("i", 1);
      FIELD_METADATA.put(1, I_FIELD_DESC);
      NAMES_TO_IDS.put("d", 2);
      THRIFT_NAMES_TO_IDS.put("d", 2);
      FIELD_METADATA.put(2, D_FIELD_DESC);
      NAMES_TO_IDS.put("s", 3);
      THRIFT_NAMES_TO_IDS.put("s", 3);
      FIELD_METADATA.put(3, S_FIELD_DESC);
      NAMES_TO_IDS.put("u", 4);
      THRIFT_NAMES_TO_IDS.put("u", 4);
      FIELD_METADATA.put(4, U_FIELD_DESC);
    }

    private java.lang.Object value;
    private short id;

    public enum TypeEnum {
      __FBTHRIFT_EMPTY_UNION__,
      I,
      D,
      S,
      U,
    }

    public static Union2 from(int _id, java.lang.Object _field) {
        return from((short) _id, _field);
    }

    public static Union2 from(short _id, java.lang.Object _field) {
        java.util.Objects.requireNonNull(_field);
        if (!FIELD_METADATA.containsKey(Integer.valueOf(_id))) {
            throw new java.lang.IllegalArgumentException("unknown field " + _id);
        }

        Union2 _u = new  Union2();

        try {
            switch(_id) {
                case 1:
                    _u.id = _id;
                    _u.value = (int) _field;
                    return _u;
                case 2:
                    _u.id = _id;
                    _u.value = (double) _field;
                    return _u;
                case 3:
                    _u.id = _id;
                    _u.value = (test.fixtures.constants.Struct1) _field;
                    return _u;
                case 4:
                    _u.id = _id;
                    _u.value = (test.fixtures.constants.Union1) _field;
                    return _u;
                default:
                throw new IllegalArgumentException("invalid type " + _field.getClass().getName() + " for field " + _id);
            }
        } catch (java.lang.Exception t) {
            throw new IllegalArgumentException("invalid type " + _field.getClass().getName() + " for field " + _id);
        }
    }

    @ThriftConstructor
    public Union2() {
    }
    
    @ThriftConstructor
    @Deprecated
    public Union2(final int i) {
        this.value = i;
        this.id = 1;
    }
    
    @ThriftConstructor
    @Deprecated
    public Union2(final double d) {
        this.value = d;
        this.id = 2;
    }
    
    @ThriftConstructor
    @Deprecated
    public Union2(final test.fixtures.constants.Struct1 s) {
        if (!Union2.allowNullFieldValues && s == null) {
            throw new TProtocolException("Cannot initialize Union field 'Union2.s' with null value!");
        }
        this.value = s;
        this.id = 3;
    }
    
    @ThriftConstructor
    @Deprecated
    public Union2(final test.fixtures.constants.Union1 u) {
        if (!Union2.allowNullFieldValues && u == null) {
            throw new TProtocolException("Cannot initialize Union field 'Union2.u' with null value!");
        }
        this.value = u;
        this.id = 4;
    }
    
    public static Union2 fromI(final int i) {
        Union2 res = new Union2();
        res.value = i;
        res.id = 1;
        return res;
    }
    
    public static Union2 fromD(final double d) {
        Union2 res = new Union2();
        res.value = d;
        res.id = 2;
        return res;
    }
    
    public static Union2 fromS(final test.fixtures.constants.Struct1 s) {
        Union2 res = new Union2();
        if (!Union2.allowNullFieldValues && s == null) {
            throw new TProtocolException("Cannot initialize Union field 'Union2.s' with null value!");
        }
        res.value = s;
        res.id = 3;
        return res;
    }
    
    public static Union2 fromU(final test.fixtures.constants.Union1 u) {
        Union2 res = new Union2();
        if (!Union2.allowNullFieldValues && u == null) {
            throw new TProtocolException("Cannot initialize Union field 'Union2.u' with null value!");
        }
        res.value = u;
        res.id = 4;
        return res;
    }

    

    @com.facebook.swift.codec.ThriftField(value=1, name="i", requiredness=Requiredness.NONE)
    public int getI() {
        if (this.id != 1) {
            throw new IllegalStateException("Not a i element!");
        }
        return (int) value;
    }

    public boolean isSetI() {
        return this.id == 1;
    }

    @com.facebook.swift.codec.ThriftField(value=2, name="d", requiredness=Requiredness.NONE)
    public double getD() {
        if (this.id != 2) {
            throw new IllegalStateException("Not a d element!");
        }
        return (double) value;
    }

    public boolean isSetD() {
        return this.id == 2;
    }

    @com.facebook.swift.codec.ThriftField(value=3, name="s", requiredness=Requiredness.NONE)
    public test.fixtures.constants.Struct1 getS() {
        if (this.id != 3) {
            throw new IllegalStateException("Not a s element!");
        }
        return (test.fixtures.constants.Struct1) value;
    }

    public boolean isSetS() {
        return this.id == 3;
    }

    @com.facebook.swift.codec.ThriftField(value=4, name="u", requiredness=Requiredness.NONE)
    public test.fixtures.constants.Union1 getU() {
        if (this.id != 4) {
            throw new IllegalStateException("Not a u element!");
        }
        return (test.fixtures.constants.Union1) value;
    }

    public boolean isSetU() {
        return this.id == 4;
    }

    @ThriftUnionId
    public short getThriftId() {
        return this.id;
    }

    public TypeEnum getThriftUnionType() {
      switch(this.id) {
        case 0:
          return TypeEnum.__FBTHRIFT_EMPTY_UNION__;
        case 1:
          return TypeEnum.I;
        case 2:
          return TypeEnum.D;
        case 3:
          return TypeEnum.S;
        case 4:
          return TypeEnum.U;
        default:
          throw new IllegalStateException("unreachable");
      }
    }

    public String getThriftName() {
        TField tField = (TField) FIELD_METADATA.get((int) this.id);
        if (tField == null) {
            return "null";
        } else {
            return tField.name;
        }
    }

    public <T> T accept(Visitor<T> visitor) {
        if (isSetI()) {
            return visitor.visitI(getI());
        }
        if (isSetD()) {
            return visitor.visitD(getD());
        }
        if (isSetS()) {
            return visitor.visitS(getS());
        }
        if (isSetU()) {
            return visitor.visitU(getU());
        }

        throw new IllegalStateException("Visitor missing for type " + this.getThriftUnionType());
    }

    @java.lang.Override
    public String toString() {
        return toStringHelper(this)
            .add("value", value)
            .add("id", id)
            .add("name", getThriftName())
            .add("type", value == null ? "<null>" : value.getClass().getSimpleName())
            .toString();
    }

    @java.lang.Override
    public boolean equals(java.lang.Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }

        Union2 other = (Union2)o;

        return Objects.equals(this.id, other.id)
                && Objects.deepEquals(this.value, other.value);
    }

    @java.lang.Override
    public int hashCode() {
        return Arrays.deepHashCode(new java.lang.Object[] {
            id,
            value,
        });
    }

    public interface Visitor<T> {
        default T visit(Union2 acceptor) {
        return acceptor.accept(this);
        }

        T visitI(int i);
        T visitD(double d);
        T visitS(test.fixtures.constants.Struct1 s);
        T visitU(test.fixtures.constants.Union1 u);
    }

    public void write0(TProtocol oprot) throws TException {
      if (this.id != 0 && this.value == null ){
        if(allowNullFieldValues) {
          // Warning: this path will generate corrupt serialized data!
          return;
        } else {
          throw new TProtocolException("Cannot write a Union with marked-as-set but null value!");
        }
      }
      oprot.writeStructBegin(STRUCT_DESC);
      switch (this.id) {
      case _I: {
        oprot.writeFieldBegin(I_FIELD_DESC);
        int i = (int)this.value;
        oprot.writeI32(i);
        oprot.writeFieldEnd();
        break;
      }
      case _D: {
        oprot.writeFieldBegin(D_FIELD_DESC);
        double d = (double)this.value;
        oprot.writeDouble(d);
        oprot.writeFieldEnd();
        break;
      }
      case _S: {
        oprot.writeFieldBegin(S_FIELD_DESC);
        test.fixtures.constants.Struct1 s = (test.fixtures.constants.Struct1)this.value;
        s.write0(oprot);
        oprot.writeFieldEnd();
        break;
      }
      case _U: {
        oprot.writeFieldBegin(U_FIELD_DESC);
        test.fixtures.constants.Union1 u = (test.fixtures.constants.Union1)this.value;
        u.write0(oprot);
        oprot.writeFieldEnd();
        break;
      }
      default:
          // ignore unknown field
      }
      oprot.writeFieldStop();
      oprot.writeStructEnd();
    }

    
    public static com.facebook.thrift.payload.Reader<Union2> asReader() {
      return Union2::read0;
    }
    
    public static Union2 read0(TProtocol oprot) throws TException {
      Union2 res = new Union2();
      res.value = null;
      res.id = (short) 0;
      oprot.readStructBegin(Union2.NAMES_TO_IDS, Union2.THRIFT_NAMES_TO_IDS, Union2.FIELD_METADATA);
      TField __field = oprot.readFieldBegin();
      if (__field.type != TType.STOP) {
          switch (__field.id) {
          case _I:
            if (__field.type == I_FIELD_DESC.type) {
              int i = oprot.readI32();
              res.value = i;
            }
            break;
          case _D:
            if (__field.type == D_FIELD_DESC.type) {
              double d = oprot.readDouble();
              res.value = d;
            }
            break;
          case _S:
            if (__field.type == S_FIELD_DESC.type) {
              test.fixtures.constants.Struct1 s = test.fixtures.constants.Struct1.read0(oprot);
              res.value = s;
            }
            break;
          case _U:
            if (__field.type == U_FIELD_DESC.type) {
              test.fixtures.constants.Union1 u = test.fixtures.constants.Union1.read0(oprot);
              res.value = u;
            }
            break;
          default:
            TProtocolUtil.skip(oprot, __field.type);
          }
        if (res.value != null) {
          res.id = __field.id;
        }
        oprot.readFieldEnd();
        TField __stopField = oprot.readFieldBegin(); // Consume the STOP byte
        if (__stopField.type != TType.STOP) {
          throw new TProtocolException(TProtocolException.INVALID_DATA, "Union 'Union2' is missing a STOP byte");
        }
      }
      oprot.readStructEnd();
      return res;
    }
    public static Union2 defaultInstance() {
        return _DEFAULT;
    }

}
