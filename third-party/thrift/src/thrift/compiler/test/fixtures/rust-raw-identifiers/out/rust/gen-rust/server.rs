// @generated by Thrift for thrift/compiler/test/fixtures/rust-raw-identifiers/src/mod.thrift
// This file is probably not the place you want to edit!

//! Server definitions for `mod`.

#![recursion_limit = "100000000"]
#![allow(non_camel_case_types, non_snake_case, non_upper_case_globals, unused_crate_dependencies, unused_imports, clippy::all)]


#[doc(inline)]
pub use :: as types;

pub mod errors {
    #[doc(inline)]
    pub use ::::services::foo;
    #[doc(inline)]
    #[allow(ambiguous_glob_reexports)]
    pub use ::::services::foo::*;
}

pub(crate) use crate as server;
pub(crate) use ::::services;



#[::async_trait::async_trait]
pub trait Foo: ::std::marker::Send + ::std::marker::Sync + 'static {
    async fn r#return(
        &self,
        _bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::ReturnExn> {
        ::std::result::Result::Err(crate::services::foo::ReturnExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "Foo",
                "return",
            ),
        ))
    }
    async fn super_(
        &self,
        _bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::SuperExn> {
        ::std::result::Result::Err(crate::services::foo::SuperExn::ApplicationException(
            ::fbthrift::ApplicationException::unimplemented_method(
                "Foo",
                "super",
            ),
        ))
    }
}

#[::async_trait::async_trait]
impl<T> Foo for ::std::boxed::Box<T>
where
    T: Foo + Send + Sync + ?Sized,
{
    async fn r#return(
        &self,
        bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::ReturnExn> {
        (**self).r#return(
            bar,
        ).await
    }
    async fn super_(
        &self,
        bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::SuperExn> {
        (**self).super_(
            bar,
        ).await
    }
}

#[::async_trait::async_trait]
impl<T> Foo for ::std::sync::Arc<T>
where
    T: Foo + Send + Sync + ?Sized,
{
    async fn r#return(
        &self,
        bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::ReturnExn> {
        (**self).r#return(
            bar,
        ).await
    }
    async fn super_(
        &self,
        bar: crate::types::ThereAreNoPascalCaseKeywords,
    ) -> ::std::result::Result<(), crate::services::foo::SuperExn> {
        (**self).super_(
            bar,
        ).await
    }
}
/// Processor for Foo's methods.
#[derive(Clone, Debug)]
pub struct FooProcessor<P, H, R, RS> {
    service: H,
    supa: ::fbthrift::NullServiceProcessor<P, R, RS>,
    _phantom: ::std::marker::PhantomData<(P, H, R, RS)>,
}


struct Args_Foo_return {
    bar: crate::types::ThereAreNoPascalCaseKeywords,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_Foo_return {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "Foo.return"))]
    fn rs_thrift_read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("bar", ::fbthrift::TType::Struct, 1),
        ];
        let mut field_bar = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::Struct, 1) => field_bar = ::std::option::Option::Some(::anyhow::Context::context(::fbthrift::Deserialize::rs_thrift_read(p), ::fbthrift::errors::DeserializingArgError { arg: "bar", function: "return"})?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            bar: field_bar.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "Foo.return", "bar"))?,
        })
    }
}


struct Args_Foo_super {
    bar: crate::types::ThereAreNoPascalCaseKeywords,
}

impl<P: ::fbthrift::ProtocolReader> ::fbthrift::Deserialize<P> for self::Args_Foo_super {
    #[inline]
    #[::tracing::instrument(skip_all, level = "trace", name = "deserialize_args", fields(method = "Foo.super"))]
    fn rs_thrift_read(p: &mut P) -> ::anyhow::Result<Self> {
        static ARGS: &[::fbthrift::Field] = &[
            ::fbthrift::Field::new("bar", ::fbthrift::TType::Struct, 1),
        ];
        let mut field_bar = ::std::option::Option::None;
        let _ = p.read_struct_begin(|_| ())?;
        loop {
            let (_, fty, fid) = p.read_field_begin(|_| (), ARGS)?;
            match (fty, fid as ::std::primitive::i32) {
                (::fbthrift::TType::Stop, _) => break,
                (::fbthrift::TType::Struct, 1) => field_bar = ::std::option::Option::Some(::anyhow::Context::context(::fbthrift::Deserialize::rs_thrift_read(p), ::fbthrift::errors::DeserializingArgError { arg: "bar", function: "super"})?),
                (fty, _) => p.skip(fty)?,
            }
            p.read_field_end()?;
        }
        p.read_struct_end()?;
        ::std::result::Result::Ok(Self {
            bar: field_bar.ok_or_else(|| ::anyhow::anyhow!("`{}` missing arg `{}`", "Foo.super", "bar"))?,
        })
    }
}

impl<P, H, R, RS> FooProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Frame: ::std::marker::Send + 'static,
    P::Deserializer: ::std::marker::Send,
    H: Foo,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    pub fn new(service: H) -> Self {
        Self {
            service,
            supa: ::fbthrift::NullServiceProcessor::new(),
            _phantom: ::std::marker::PhantomData,
        }
    }

    pub fn into_inner(self) -> H {
        self.service
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "Foo.return"))]
    async fn handle_return<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"Foo";
        const METHOD_NAME: &::std::ffi::CStr = c"return";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"Foo.return";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_Foo_return = ::fbthrift::Deserialize::rs_thrift_read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.r#return(
                _args.bar,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "Foo.return", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "Foo.return", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("Foo.return", exn);
                ::tracing::error!(method = "Foo.return", panic = ?aexn);
                ::std::result::Result::Err(crate::services::foo::ReturnExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::foo::ReturnExn>(
            "return",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }

    #[::tracing::instrument(skip_all, name = "handler", fields(method = "Foo.super"))]
    async fn handle_super<'a>(
        &'a self,
        p: &'a mut P::Deserializer,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        use ::futures::FutureExt as _;

        const SERVICE_NAME: &::std::ffi::CStr = c"Foo";
        const METHOD_NAME: &::std::ffi::CStr = c"super";
        const SERVICE_METHOD_NAME: &::std::ffi::CStr = c"Foo.super";
        let mut ctx_stack = req_ctxt.get_context_stack(SERVICE_NAME, SERVICE_METHOD_NAME)?;
        ::fbthrift::ContextStack::pre_read(&mut ctx_stack)?;
        let _args: self::Args_Foo_super = ::fbthrift::Deserialize::rs_thrift_read(p)?;
        let bytes_read = ::fbthrift::help::buf_len(&req)?;
        ::fbthrift::ContextStack::on_read_data(&mut ctx_stack, ::fbthrift::SerializedMessage {
            protocol: P::PROTOCOL_ID,
            method_name: METHOD_NAME,
            buffer: req,
        })?;
        ::fbthrift::ContextStack::post_read(&mut ctx_stack, bytes_read)?;

        let res = ::std::panic::AssertUnwindSafe(
            self.service.super_(
                _args.bar,
            )
        )
        .catch_unwind()
        .await;

        // nested results - panic catch on the outside, method on the inside
        let res = match res {
            ::std::result::Result::Ok(::std::result::Result::Ok(res)) => {
                ::tracing::trace!(method = "Foo.super", "success");
                ::std::result::Result::Ok(res)
            }
            ::std::result::Result::Ok(::std::result::Result::Err(exn)) => {
                ::tracing::error!(method = "Foo.super", exception = ?exn);
                ::std::result::Result::Err(exn)
            }
            ::std::result::Result::Err(exn) => {
                let aexn = ::fbthrift::ApplicationException::handler_panic("Foo.super", exn);
                ::tracing::error!(method = "Foo.super", panic = ?aexn);
                ::std::result::Result::Err(crate::services::foo::SuperExn::ApplicationException(aexn))
            }
        };

        let env = ::fbthrift::help::serialize_result_envelope::<P, R, crate::services::foo::SuperExn>(
            "super",
            METHOD_NAME,
            _seqid,
            req_ctxt,
            &mut ctx_stack,
            res,
        )?;
        reply_state.send_reply(env);
        ::std::result::Result::Ok(())
    }
}

#[::async_trait::async_trait]
impl<P, H, R, RS> ::fbthrift::ServiceProcessor<P> for FooProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Deserializer: ::std::marker::Send,
    H: Foo,
    P::Frame: ::std::marker::Send + 'static,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    type RequestContext = R;
    type ReplyState = RS;

    #[inline]
    fn method_idx(&self, name: &[::std::primitive::u8]) -> ::std::result::Result<::std::primitive::usize, ::fbthrift::ApplicationException> {
        match name {
            b"return" => ::std::result::Result::Ok(0usize),
            b"super" => ::std::result::Result::Ok(1usize),
            _ => ::std::result::Result::Err(::fbthrift::ApplicationException::unknown_method()),
        }
    }

    #[allow(clippy::match_single_binding)]
    async fn handle_method(
        &self,
        idx: ::std::primitive::usize,
        _p: &mut P::Deserializer,
        _req: ::fbthrift::ProtocolDecoded<P>,
        _req_ctxt: &R,
        _reply_state: ::std::sync::Arc<RS>,
        _seqid: ::std::primitive::u32,
    ) -> ::anyhow::Result<()> {
        match idx {
            0usize => {
                self.handle_return(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            1usize => {
                self.handle_super(_p, _req, _req_ctxt, _reply_state, _seqid).await
            }
            bad => panic!(
                "{}: unexpected method idx {}",
                "FooProcessor",
                bad
            ),
        }
    }

    #[allow(clippy::match_single_binding)]
    #[inline]
    fn create_interaction_idx(&self, name: &::std::primitive::str) -> ::anyhow::Result<::std::primitive::usize> {
        match name {
            _ => ::anyhow::bail!("Unknown interaction"),
        }
    }

    #[allow(clippy::match_single_binding)]
    fn handle_create_interaction(
        &self,
        idx: ::std::primitive::usize,
    ) -> ::anyhow::Result<
        ::std::sync::Arc<dyn ::fbthrift::ThriftService<P::Frame, Handler = (), RequestContext = Self::RequestContext, ReplyState = Self::ReplyState> + ::std::marker::Send + 'static>
    > {
        match idx {
            bad => panic!(
                "{}: unexpected method idx {}",
                "FooProcessor",
                bad
            ),
        }
    }

    async fn handle_on_termination(&self) {
    }
}

#[::async_trait::async_trait]
impl<P, H, R, RS> ::fbthrift::ThriftService<P::Frame> for FooProcessor<P, H, R, RS>
where
    P: ::fbthrift::Protocol + ::std::marker::Send + ::std::marker::Sync + 'static,
    P::Deserializer: ::std::marker::Send,
    P::Frame: ::std::marker::Send + 'static,
    H: Foo,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = <P as ::fbthrift::Protocol>::Frame>
        + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<P::Frame, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::ProtocolDecoded<P>: ::std::clone::Clone,
    ::fbthrift::ProtocolEncodedFinal<P>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    type Handler = H;
    type RequestContext = R;
    type ReplyState = RS;

    #[tracing::instrument(level="trace", skip_all, fields(service = "Foo"))]
    async fn call(
        &self,
        req: ::fbthrift::ProtocolDecoded<P>,
        req_ctxt: &R,
        reply_state: ::std::sync::Arc<RS>,
    ) -> ::anyhow::Result<()> {
        use ::fbthrift::{ProtocolReader as _, ServiceProcessor as _};
        let mut p = P::deserializer(req.clone());
        let (idx, mty, seqid) = p.read_message_begin(|name| self.method_idx(name))?;
        if mty != ::fbthrift::MessageType::Call {
            return ::std::result::Result::Err(::std::convert::From::from(::fbthrift::ApplicationException::new(
                ::fbthrift::ApplicationExceptionErrorCode::InvalidMessageType,
                format!("message type {:?} not handled", mty)
            )));
        }
        let idx = match idx {
            ::std::result::Result::Ok(idx) => idx,
            ::std::result::Result::Err(_) => {
                return self.supa.call(req, req_ctxt, reply_state).await;
            }
        };
        self.handle_method(idx, &mut p, req, req_ctxt, reply_state, seqid).await?;
        p.read_message_end()?;

        ::std::result::Result::Ok(())
    }

    fn create_interaction(
        &self,
        name: &::std::primitive::str,
    ) -> ::anyhow::Result<
        ::std::sync::Arc<dyn ::fbthrift::ThriftService<P::Frame, Handler = (), RequestContext = R, ReplyState = RS> + ::std::marker::Send + 'static>
    > {
        use ::fbthrift::{ServiceProcessor as _};
        let idx = self.create_interaction_idx(name);
        let idx = match idx {
            ::anyhow::Result::Ok(idx) => idx,
            ::anyhow::Result::Err(_) => {
                return self.supa.create_interaction(name);
            }
        };
        self.handle_create_interaction(idx)
    }

    fn get_method_metadata(&self) -> &'static [::fbthrift::processor::MethodMetadata] {
        &[
            // From mod.Foo:
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "return",
              starts_interaction: false,
              interaction_name: None,
            },
            ::fbthrift::processor::MethodMetadata{
              interaction_type: ::fbthrift::processor::InteractionType::None,
              rpc_kind: ::fbthrift::processor::RpcKind::SINGLE_REQUEST_SINGLE_RESPONSE,
              name: "super",
              starts_interaction: false,
              interaction_name: None,
            },
        ]
    }

    async fn on_termination(&self) {
        use ::fbthrift::{ServiceProcessor as _};
        self.handle_on_termination().await
    }
}

/// Construct a new instance of a Foo service.
///
/// This is called when a new instance of a Thrift service Processor
/// is needed for a particular Thrift protocol.
#[::tracing::instrument(level="debug", skip_all, fields(proto = ?proto))]
pub fn make_Foo_server<F, H, R, RS>(
    proto: ::fbthrift::ProtocolID,
    handler: H,
) -> ::std::result::Result<::std::boxed::Box<dyn ::fbthrift::ThriftService<F, Handler = H, RequestContext = R, ReplyState = RS> + ::std::marker::Send + 'static>, ::fbthrift::ApplicationException>
where
    F: ::fbthrift::Framing + ::std::marker::Send + ::std::marker::Sync + 'static,
    H: Foo,
    R: ::fbthrift::RequestContext<Name = ::std::ffi::CStr> + ::std::marker::Send + ::std::marker::Sync + 'static,
    <R as ::fbthrift::RequestContext>::ContextStack: ::fbthrift::ContextStack<Name = R::Name, Frame = F> + ::std::marker::Send + ::std::marker::Sync + 'static,
    RS: ::fbthrift::ReplyState<F, RequestContext = R> + ::std::marker::Send + ::std::marker::Sync + 'static,
    ::fbthrift::FramingDecoded<F>: ::std::clone::Clone,
    ::fbthrift::FramingEncodedFinal<F>: ::std::clone::Clone + ::fbthrift::BufExt,
{
    match proto {
        ::fbthrift::ProtocolID::BinaryProtocol => {
            ::std::result::Result::Ok(::std::boxed::Box::new(FooProcessor::<::fbthrift::BinaryProtocol<F>, H, R, RS>::new(handler)))
        }
        ::fbthrift::ProtocolID::CompactProtocol => {
            ::std::result::Result::Ok(::std::boxed::Box::new(FooProcessor::<::fbthrift::CompactProtocol<F>, H, R, RS>::new(handler)))
        }
        bad => {
            ::tracing::error!(method = "Foo.", invalid_protocol = ?bad);
            ::std::result::Result::Err(::fbthrift::ApplicationException::invalid_protocol(bad))
        }
    }
}

