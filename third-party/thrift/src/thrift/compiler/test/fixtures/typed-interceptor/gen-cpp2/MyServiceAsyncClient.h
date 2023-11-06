/**
 * Autogenerated by Thrift for thrift/compiler/test/fixtures/typed-interceptor/src/module.thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated @nocommit
 */
#pragma once

#include <thrift/lib/cpp2/gen/client_h.h>

#include "thrift/compiler/test/fixtures/typed-interceptor/gen-cpp2/module_types.h"

namespace apache { namespace thrift {
  class Cpp2RequestContext;
  namespace detail { namespace ac { struct ClientRequestContext; }}
  namespace transport { class THeader; }
}}

namespace cpp2 {
class MyService;
} // cpp2
namespace apache::thrift {

template <>
class Client<::cpp2::MyService> : public apache::thrift::GeneratedAsyncClient {
 public:
  using apache::thrift::GeneratedAsyncClient::GeneratedAsyncClient;

  char const* getServiceName() const noexcept override {
    return "MyService";
  }


  virtual void echo(std::unique_ptr<apache::thrift::RequestCallback> callback);
  virtual void echo(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback);
 protected:
  void echoImpl(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, bool stealRpcOptions = false);
 public:

  virtual void sync_echo();
  virtual void sync_echo(apache::thrift::RpcOptions& rpcOptions);

  virtual folly::Future<folly::Unit> future_echo();
  virtual folly::SemiFuture<folly::Unit> semifuture_echo();
  virtual folly::Future<folly::Unit> future_echo(apache::thrift::RpcOptions& rpcOptions);
  virtual folly::SemiFuture<folly::Unit> semifuture_echo(apache::thrift::RpcOptions& rpcOptions);
  virtual folly::Future<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> header_future_echo(apache::thrift::RpcOptions& rpcOptions);
  virtual folly::SemiFuture<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> header_semifuture_echo(apache::thrift::RpcOptions& rpcOptions);

#if FOLLY_HAS_COROUTINES
#if __clang__
  template <int = 0>
  folly::coro::Task<void> co_echo() {
    return co_echo<false>(nullptr);
  }
  template <int = 0>
  folly::coro::Task<void> co_echo(apache::thrift::RpcOptions& rpcOptions) {
    return co_echo<true>(&rpcOptions);
  }
#else
  folly::coro::Task<void> co_echo() {
    co_await folly::coro::detachOnCancel(semifuture_echo());
  }
  folly::coro::Task<void> co_echo(apache::thrift::RpcOptions& rpcOptions) {
    co_await folly::coro::detachOnCancel(semifuture_echo(rpcOptions));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<void> co_echo(apache::thrift::RpcOptions* rpcOptions) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = echoCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if constexpr (hasRpcOptions) {
      echoImpl(*rpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback));
    } else {
      echoImpl(*defaultRpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback));
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    if (auto ew = recv_wrapped_echo(returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  virtual void echo(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback);


  static folly::exception_wrapper recv_wrapped_echo(::apache::thrift::ClientReceiveState& state);
  static void recv_echo(::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  virtual void recv_instance_echo(::apache::thrift::ClientReceiveState& state);
  virtual folly::exception_wrapper recv_instance_wrapped_echo(::apache::thrift::ClientReceiveState& state);
 private:
  template <typename Protocol_, typename RpcOptions>
  void echoT(Protocol_* prot, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> echoCtx(apache::thrift::RpcOptions* rpcOptions);
 public:
  virtual void getRandomData(std::unique_ptr<apache::thrift::RequestCallback> callback, const ::cpp2::Request& p_req);
  virtual void getRandomData(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, const ::cpp2::Request& p_req);
 protected:
  void getRandomDataImpl(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::cpp2::Request& p_req, bool stealRpcOptions = false);
 public:

  virtual void sync_getRandomData(::std::string& _return, const ::cpp2::Request& p_req);
  virtual void sync_getRandomData(apache::thrift::RpcOptions& rpcOptions, ::std::string& _return, const ::cpp2::Request& p_req);

  virtual folly::Future<::std::string> future_getRandomData(const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<::std::string> semifuture_getRandomData(const ::cpp2::Request& p_req);
  virtual folly::Future<::std::string> future_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<::std::string> semifuture_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::Future<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> header_future_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<std::pair<::std::string, std::unique_ptr<apache::thrift::transport::THeader>>> header_semifuture_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);

#if FOLLY_HAS_COROUTINES
#if __clang__
  template <int = 0>
  folly::coro::Task<::std::string> co_getRandomData(const ::cpp2::Request& p_req) {
    return co_getRandomData<false>(nullptr, p_req);
  }
  template <int = 0>
  folly::coro::Task<::std::string> co_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req) {
    return co_getRandomData<true>(&rpcOptions, p_req);
  }
#else
  folly::coro::Task<::std::string> co_getRandomData(const ::cpp2::Request& p_req) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getRandomData(p_req));
  }
  folly::coro::Task<::std::string> co_getRandomData(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getRandomData(rpcOptions, p_req));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<::std::string> co_getRandomData(apache::thrift::RpcOptions* rpcOptions, const ::cpp2::Request& p_req) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = getRandomDataCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if constexpr (hasRpcOptions) {
      getRandomDataImpl(*rpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_req);
    } else {
      getRandomDataImpl(*defaultRpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_req);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    ::std::string _return;
    if (auto ew = recv_wrapped_getRandomData(_return, returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
    co_return _return;
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  virtual void getRandomData(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, const ::cpp2::Request& p_req);


  static folly::exception_wrapper recv_wrapped_getRandomData(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  static void recv_getRandomData(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  virtual void recv_instance_getRandomData(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
  virtual folly::exception_wrapper recv_instance_wrapped_getRandomData(::std::string& _return, ::apache::thrift::ClientReceiveState& state);
 private:
  template <typename Protocol_, typename RpcOptions>
  void getRandomDataT(Protocol_* prot, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::cpp2::Request& p_req);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> getRandomDataCtx(apache::thrift::RpcOptions* rpcOptions);
 public:
  virtual void getId(std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int32_t p_field);
  virtual void getId(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, ::std::int32_t p_field);
 protected:
  void getIdImpl(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int32_t p_field, bool stealRpcOptions = false);
 public:

  virtual ::std::int32_t sync_getId(::std::int32_t p_field);
  virtual ::std::int32_t sync_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field);

  virtual folly::Future<::std::int32_t> future_getId(::std::int32_t p_field);
  virtual folly::SemiFuture<::std::int32_t> semifuture_getId(::std::int32_t p_field);
  virtual folly::Future<::std::int32_t> future_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field);
  virtual folly::SemiFuture<::std::int32_t> semifuture_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field);
  virtual folly::Future<std::pair<::std::int32_t, std::unique_ptr<apache::thrift::transport::THeader>>> header_future_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field);
  virtual folly::SemiFuture<std::pair<::std::int32_t, std::unique_ptr<apache::thrift::transport::THeader>>> header_semifuture_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field);

#if FOLLY_HAS_COROUTINES
#if __clang__
  template <int = 0>
  folly::coro::Task<::std::int32_t> co_getId(::std::int32_t p_field) {
    return co_getId<false>(nullptr, p_field);
  }
  template <int = 0>
  folly::coro::Task<::std::int32_t> co_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field) {
    return co_getId<true>(&rpcOptions, p_field);
  }
#else
  folly::coro::Task<::std::int32_t> co_getId(::std::int32_t p_field) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getId(p_field));
  }
  folly::coro::Task<::std::int32_t> co_getId(apache::thrift::RpcOptions& rpcOptions, ::std::int32_t p_field) {
    co_return co_await folly::coro::detachOnCancel(semifuture_getId(rpcOptions, p_field));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<::std::int32_t> co_getId(apache::thrift::RpcOptions* rpcOptions, ::std::int32_t p_field) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = getIdCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if constexpr (hasRpcOptions) {
      getIdImpl(*rpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_field);
    } else {
      getIdImpl(*defaultRpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_field);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    ::std::int32_t _return;
    if (auto ew = recv_wrapped_getId(_return, returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
    co_return _return;
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  virtual void getId(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, ::std::int32_t p_field);


  static folly::exception_wrapper recv_wrapped_getId(::std::int32_t& _return, ::apache::thrift::ClientReceiveState& state);
  static ::std::int32_t recv_getId(::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  virtual ::std::int32_t recv_instance_getId(::apache::thrift::ClientReceiveState& state);
  virtual folly::exception_wrapper recv_instance_wrapped_getId(::std::int32_t& _return, ::apache::thrift::ClientReceiveState& state);
 private:
  template <typename Protocol_, typename RpcOptions>
  void getIdT(Protocol_* prot, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, ::std::int32_t p_field);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> getIdCtx(apache::thrift::RpcOptions* rpcOptions);
 public:
  virtual void ping_eb(std::unique_ptr<apache::thrift::RequestCallback> callback, const ::cpp2::Request& p_req);
  virtual void ping_eb(apache::thrift::RpcOptions& rpcOptions, std::unique_ptr<apache::thrift::RequestCallback> callback, const ::cpp2::Request& p_req);
 protected:
  void ping_ebImpl(apache::thrift::RpcOptions& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::cpp2::Request& p_req, bool stealRpcOptions = false);
 public:

  virtual void sync_ping_eb(const ::cpp2::Request& p_req);
  virtual void sync_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);

  virtual folly::Future<folly::Unit> future_ping_eb(const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<folly::Unit> semifuture_ping_eb(const ::cpp2::Request& p_req);
  virtual folly::Future<folly::Unit> future_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<folly::Unit> semifuture_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::Future<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> header_future_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);
  virtual folly::SemiFuture<std::pair<folly::Unit, std::unique_ptr<apache::thrift::transport::THeader>>> header_semifuture_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req);

#if FOLLY_HAS_COROUTINES
#if __clang__
  template <int = 0>
  folly::coro::Task<void> co_ping_eb(const ::cpp2::Request& p_req) {
    return co_ping_eb<false>(nullptr, p_req);
  }
  template <int = 0>
  folly::coro::Task<void> co_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req) {
    return co_ping_eb<true>(&rpcOptions, p_req);
  }
#else
  folly::coro::Task<void> co_ping_eb(const ::cpp2::Request& p_req) {
    co_await folly::coro::detachOnCancel(semifuture_ping_eb(p_req));
  }
  folly::coro::Task<void> co_ping_eb(apache::thrift::RpcOptions& rpcOptions, const ::cpp2::Request& p_req) {
    co_await folly::coro::detachOnCancel(semifuture_ping_eb(rpcOptions, p_req));
  }
#endif
 private:
  template <bool hasRpcOptions>
  folly::coro::Task<void> co_ping_eb(apache::thrift::RpcOptions* rpcOptions, const ::cpp2::Request& p_req) {
    const folly::CancellationToken& cancelToken =
        co_await folly::coro::co_current_cancellation_token;
    const bool cancellable = cancelToken.canBeCancelled();
    apache::thrift::ClientReceiveState returnState;
    apache::thrift::ClientCoroCallback<false> callback(&returnState, co_await folly::coro::co_current_executor);
    auto protocolId = apache::thrift::GeneratedAsyncClient::getChannel()->getProtocolId();
    auto [ctx, header] = ping_ebCtx(rpcOptions);
    using CancellableCallback = apache::thrift::CancellableRequestClientCallback<false>;
    auto cancellableCallback = cancellable ? CancellableCallback::create(&callback, channel_) : nullptr;
    static apache::thrift::RpcOptions* defaultRpcOptions = new apache::thrift::RpcOptions();
    auto wrappedCallback = apache::thrift::RequestClientCallback::Ptr(cancellableCallback ? (apache::thrift::RequestClientCallback*)cancellableCallback.get() : &callback);
    if constexpr (hasRpcOptions) {
      ping_ebImpl(*rpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_req);
    } else {
      ping_ebImpl(*defaultRpcOptions, std::move(header), ctx.get(), std::move(wrappedCallback), p_req);
    }
    if (cancellable) {
      folly::CancellationCallback cb(cancelToken, [&] { CancellableCallback::cancel(std::move(cancellableCallback)); });
      co_await callback.co_waitUntilDone();
    } else {
      co_await callback.co_waitUntilDone();
    }
    if (returnState.isException()) {
      co_yield folly::coro::co_error(std::move(returnState.exception()));
    }
    returnState.resetProtocolId(protocolId);
    returnState.resetCtx(std::move(ctx));
    SCOPE_EXIT {
      if (hasRpcOptions && returnState.header()) {
        auto* rheader = returnState.header();
        if (!rheader->getHeaders().empty()) {
          rpcOptions->setReadHeaders(rheader->releaseHeaders());
        }
        rpcOptions->setRoutingData(rheader->releaseRoutingData());
      }
    };
    if (auto ew = recv_wrapped_ping_eb(returnState)) {
      co_yield folly::coro::co_error(std::move(ew));
    }
  }
 public:
#endif // FOLLY_HAS_COROUTINES

  virtual void ping_eb(folly::Function<void (::apache::thrift::ClientReceiveState&&)> callback, const ::cpp2::Request& p_req);


  static folly::exception_wrapper recv_wrapped_ping_eb(::apache::thrift::ClientReceiveState& state);
  static void recv_ping_eb(::apache::thrift::ClientReceiveState& state);
  // Mock friendly virtual instance method
  virtual void recv_instance_ping_eb(::apache::thrift::ClientReceiveState& state);
  virtual folly::exception_wrapper recv_instance_wrapped_ping_eb(::apache::thrift::ClientReceiveState& state);
 private:
  template <typename Protocol_, typename RpcOptions>
  void ping_ebT(Protocol_* prot, RpcOptions&& rpcOptions, std::shared_ptr<apache::thrift::transport::THeader> header, apache::thrift::ContextStack* contextStack, apache::thrift::RequestClientCallback::Ptr callback, const ::cpp2::Request& p_req);
  std::pair<::apache::thrift::ContextStack::UniquePtr, std::shared_ptr<::apache::thrift::transport::THeader>> ping_ebCtx(apache::thrift::RpcOptions* rpcOptions);
 public:
};

} // namespace apache::thrift

namespace cpp2 {
using MyServiceAsyncClient [[deprecated("Use apache::thrift::Client<MyService> instead")]] = ::apache::thrift::Client<MyService>;
} // cpp2